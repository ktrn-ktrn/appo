package mappers

import (
	"appo/app/models/entities"
	"database/sql"
	"fmt"
)

type InterviewerMapper struct {
	db *sql.DB
}

func (m *InterviewerMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

//выбрать всех сотрудников
func (m *InterviewerMapper) SelectAll() ([]*entities.Interviewer, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_position     sql.NullString
	)

	interviewers := make([]*entities.Interviewer, 0)
	//обращение к БД
	//выбираем всех сотрудников из таблицы t_interviewer
	rows, err := m.db.Query("SELECT c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_position FROM t_interviewer")
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("InterviewerMapper::Select:%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_position)
		interviewer := &entities.Interviewer{
			ID:          c_id.Int64,
			Surname:     c_surname.String,
			Name:        c_name.String,
			Patronymic:  c_patronymic.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Position:    c_position.String,
		}
		interviewers = append(interviewers, interviewer)
	}
	return interviewers, nil
}

//получить сотрудников в выбранном ассессменте
func (m *InterviewerMapper) Select(assessmentId int64) ([]*entities.Interviewer, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_position     sql.NullString
	)

	interviewers := make([]*entities.Interviewer, 0)
	rows, err := m.db.Query("SELECT u.c_id, u.c_surname, u.c_name, u.c_patronymic, c_email, c_phone_number, c_position FROM t_interviewer u INNER JOIN toc_assessment_interviewer d ON u.c_id = d.c_id_interviewer WHERE d.c_id_assessment = $1", assessmentId)
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("InterviewerMapper::Select:%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_position)
		interviewer := &entities.Interviewer{
			ID:          c_id.Int64,
			Surname:     c_surname.String,
			Name:        c_name.String,
			Patronymic:  c_patronymic.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Position:    c_position.String,
		}
		interviewers = append(interviewers, interviewer)
	}
	return interviewers, nil
}

//получить выбранного сотрудника
func (m *InterviewerMapper) SelectById(interviewerId int64) (*entities.Interviewer, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_position     sql.NullString
	)

	row := m.db.QueryRow("SELECT c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_position FROM t_interviewer WHERE c_id = $1", interviewerId)

	err := row.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_position)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("InterviewerMapper::SelectById:%v", err)
	} else if err != nil {
		return nil, fmt.Errorf("InterviewerMapper::SelectById:%v", err)
	}
	interviewer := &entities.Interviewer{
		ID:          c_id.Int64,
		Surname:     c_surname.String,
		Name:        c_name.String,
		Patronymic:  c_patronymic.String,
		Email:       c_email.String,
		PhoneNumber: c_phone_number.String,
		Position:    c_position.String,
	}
	return interviewer, nil
}

//добавить сотрудника в ассессмент
func (m *InterviewerMapper) Insert(newInterviewer *entities.Interviewer, assessmentId int64) (int64, error) {
	var insertedId int64
	//обращения к БД
	//добавление сотрудника к списку сотрудников
	insertQuery := `INSERT INTO t_interviewer 
		(c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_position) 
		SELECT nextval('interviewer_id'), $1, $2, $3, $4, $5, $6 
		WHERE NOT EXISTS(SELECT c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_position FROM t_interviewer WHERE c_surname = $7 AND c_name = $8 AND c_patronymic = $9 AND c_position = $10)`
	_, err := m.db.Exec(insertQuery, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Email, newInterviewer.PhoneNumber, newInterviewer.Position, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки сотрудника: %v", err)
	}
	//возвращаем его ID
	row := m.db.QueryRow(`select c_id FROM t_interviewer WHERE c_surname = $1 AND c_name = $2 AND c_patronymic = $3 AND c_position = $4`, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)

	err = row.Scan(&insertedId)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}
	//добавление сотрдуника в таблицу связи toc_assessment_interviewer
	insertQueryToAssess := `INSERT INTO toc_assessment_interviewer 
		(c_id, c_id_assessment, c_id_interviewer) 
		SELECT nextval('assessment_interviewer_id'), $1, $2 
		WHERE NOT EXISTS(SELECT c_id, c_id_assessment, c_id_interviewer FROM toc_assessment_interviewer WHERE c_id_interviewer = $3 AND c_id_assessment = $4)`
	_, err = m.db.Exec(insertQueryToAssess, assessmentId, insertedId, insertedId, assessmentId)

	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки сотрудника в ассессмент: %v", err)
	}
	return insertedId, nil
}

//изменить сотрудника
func (m *InterviewerMapper) Update(newInterviewer *entities.Interviewer, interviewerId int64) (int64, error) {
	insertQuery := `UPDATE t_interviewer 
		SET c_surname = $1, c_name = $2, c_patronymic = $3, c_email = $4, c_phone_number = $5, c_position = $6
		WHERE c_id = $7`
	_, err := m.db.Exec(insertQuery, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Email, newInterviewer.PhoneNumber, newInterviewer.Position, interviewerId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка обновления сотрудника: %v", err)
	}
	return interviewerId, nil
}

//добавить сотрудника к списку сотрдуников
func (m *InterviewerMapper) InsertInterviewer(newInterviewer *entities.Interviewer) (int64, error) {
	var insertedId int64

	//добавить к списку
	insertQuery := `INSERT INTO t_interviewer 
		(c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_position) 
		SELECT nextval('interviewer_id'), $1, $2, $3, $4, $5, $6 
		WHERE NOT EXISTS(SELECT c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_position FROM t_interviewer WHERE c_surname = $7 AND c_name = $8 AND c_patronymic = $9 AND c_position = $10)`
	_, err := m.db.Exec(insertQuery, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Email, newInterviewer.PhoneNumber, newInterviewer.Position, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки сотрудника: %v", err)
	}
	//вернуть его ID
	row := m.db.QueryRow(`select c_id FROM t_interviewer WHERE c_surname = $1 AND c_name = $2 AND c_patronymic = $3 AND c_position = $4`, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)

	err = row.Scan(&insertedId)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}
	return insertedId, nil
}

//удаление сотрудника из ассессмента
func (m *InterviewerMapper) Delete(id int64, idAssessment int64) error {
	_, err := m.db.Exec("DELETE FROM toc_assessment_interviewer WHERE c_id_interviewer = $1 AND c_id_assessment = $2", id, idAssessment)
	if err == sql.ErrNoRows {
		return fmt.Errorf("InterviewerMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("InterviewerMapper::Delete:%v", err)
	}
	return nil
}

//удалить из словаря
func (m *InterviewerMapper) DeleteFromD(id int64) error {

	deleteQuery := `DELETE FROM toc_assessment_interviewer WHERE c_id_interviewer = $1;`

	_, err := m.db.Exec(deleteQuery, id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	} else if err != nil {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	}

	deleteQuery = `DELETE FROM t_interviewer WHERE c_id = $1;`

	_, err = m.db.Exec(deleteQuery, id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	} else if err != nil {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	}
	return nil
}
