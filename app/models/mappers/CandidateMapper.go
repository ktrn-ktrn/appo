package mappers

import (
	"appo/app/models/entities"
	"database/sql"
	"fmt"
)

type CandidateMapper struct {
	db *sql.DB
}

func (m *CandidateMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

//получить список всех кандидатов, которые существуют
func (m *CandidateMapper) SelectAllCandidates() ([]*entities.Candidate, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_resume       sql.NullString
		c_addres       sql.NullString
		c_education    sql.NullString
		c_birth_date   sql.NullString
	)
	//создаем пустой срез кандатов
	candidates := make([]*entities.Candidate, 0)
	//обращение к БД
	query := `SELECT c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, 
		c_resume, c_addres, to_char(c_birth_date, 'YYYY-MM-DD'), c_education 
		FROM t_candidate`
	//выполняем обращение
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectAllCandidates:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		//считываем данные
		rows.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_resume, &c_addres, &c_birth_date, &c_education)
		//создаем объект саndidate и записываем в него полученные данные
		candidate := &entities.Candidate{ID: c_id.Int64,
			Surname:     c_surname.String,
			Name:        c_name.String,
			Patronymic:  c_patronymic.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Resume:      c_resume.String,
			Addres:      c_addres.String,
			Education:   c_education.String,
			BirthDate:   c_birth_date.String,
		}
		//добавляем созданный объект к срезу
		candidates = append(candidates, candidate)
	}
	//возвращаем его
	return candidates, nil
}

//получить всех кандидатов которые участвуют в выбранном асссесменте
func (m *CandidateMapper) Select(assessmentId int64) ([]*entities.Candidate, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_resume       sql.NullString
		c_addres       sql.NullString
		c_education    sql.NullString
		c_birth_date   sql.NullString
		c_status_name  sql.NullString
	)

	candidates := make([]*entities.Candidate, 0)
	//запрос к БД
	query := `SELECT u.c_id, u.c_surname, u.c_name, u.c_patronymic, u.c_email, u.c_phone_number, 
		u.c_resume, u.c_addres, to_char(u.c_birth_date, 'DD.MM.YYYY'), u.c_education, v.c_status 
		FROM t_candidate u INNER JOIN toc_assessment_candidate d ON u.c_id = d.c_id_candidate 
		INNER JOIN t_status_candidate v ON d.c_status_candidate = v.c_id WHERE d.c_id_assessment = $1 `
	rows, err := m.db.Query(query, assessmentId)
	if err != nil {
		return nil, fmt.Errorf("CandidateMapper::Select:%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		//считываем данные и записываем в candidate
		rows.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_resume, &c_addres, &c_birth_date, &c_education, &c_status_name)
		candidate := &entities.Candidate{ID: c_id.Int64,
			Surname:     c_surname.String,
			Name:        c_name.String,
			Patronymic:  c_patronymic.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Resume:      c_resume.String,
			Addres:      c_addres.String,
			Education:   c_education.String,
			BirthDate:   c_birth_date.String,
			StatusName:  c_status_name.String,
		}
		//добавляем candidate к созданному срезу
		candidates = append(candidates, candidate)
	}
	return candidates, nil
}

//получить выбранного кандидата
func (m *CandidateMapper) SelectById(id int64, assessmentId int64) (*entities.Candidate, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_resume       sql.NullString
		c_addres       sql.NullString
		c_education    sql.NullString
		c_birth_date   sql.NullString
		c_status_name  sql.NullString
	)
	//обращаемся к БД
	query := `SELECT u.c_id, u.c_surname, u.c_name, u.c_patronymic, u.c_email, u.c_phone_number, 
		u.c_resume, u.c_addres, u.c_birth_date, u.c_education, v.c_status 
		FROM t_candidate u INNER JOIN toc_assessment_candidate d ON u.c_id = d.c_id_candidate 
		INNER JOIN t_status_candidate v ON d.c_status_candidate = v.c_id WHERE d.c_id_assessment = $1 AND u.c_id = $2`
	//выполняем
	row := m.db.QueryRow(query, assessmentId, id)
	//считываем
	err := row.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_resume, &c_addres, &c_birth_date, &c_education, &c_status_name)
	//выдаем ошибку если по результту ничего не найдено или произошла иная ошибка
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}

	//записываем данные если нет ошибок
	candidate := &entities.Candidate{ID: c_id.Int64,
		Surname:     c_surname.String,
		Name:        c_name.String,
		Patronymic:  c_patronymic.String,
		Email:       c_email.String,
		PhoneNumber: c_phone_number.String,
		Resume:      c_resume.String,
		Addres:      c_addres.String,
		Education:   c_education.String,
		BirthDate:   c_birth_date.String,
		StatusName:  c_status_name.String,
	}

	return candidate, nil
}

//получить возможные статусы
func (m *CandidateMapper) SelectStatus(candidateId int64, assessmentId int64) ([]*entities.CandidateStatus, error) {
	var (
		c_id     sql.NullInt64
		c_status sql.NullString
	)

	statuses := make([]*entities.CandidateStatus, 0)
	//запросы к БД
	//получаем родительский статус кандидата
	query := `
	SELECT c_id, c_status FROM t_status_candidate 
		WHERE c_id = (select fk_parent FROM t_status_candidate where c_id = 
		(select c_status_candidate FROM toc_assessment_candidate where 
		c_id_candidate = $1));`
	//получаем возможные статусы
	query2 := `SELECT u.c_id, u.c_status FROM t_status_candidate u INNER JOIN toc_assessment_candidate d ON 
d.c_status_candidate = u.fk_parent WHERE d.c_id_candidate = $1 AND d.c_id_assessment = $2`

	rows, err := m.db.Query(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectStatus:%v", err)
	}

	rows2, erro := m.db.Query(query2, candidateId, assessmentId)
	if erro != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectStatus:%v", erro)
	}

	defer rows.Close()
	//получаем данные
	for rows.Next() {
		rows.Scan(&c_id, &c_status)
		status := &entities.CandidateStatus{
			ID:     c_id.Int64,
			Status: c_status.String,
		}
		statuses = append(statuses, status)
	}

	defer rows2.Close()

	for rows2.Next() {
		rows2.Scan(&c_id, &c_status)
		status := &entities.CandidateStatus{
			ID:     c_id.Int64,
			Status: c_status.String,
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

//задать статус кандидата
func (m *CandidateMapper) SetStatus(newStatus *entities.CandidateStatus, statusId int64, candidateId int64) (int64, error) {
	insertQuery := `UPDATE toc_assessment_candidate SET c_status_candidate = $1 WHERE c_id_candidate = $2`
	_, err := m.db.Exec(insertQuery, statusId, candidateId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка изменения статуса кандидата: %v", err)
	}
	return statusId, nil
}

//создание кандидата
func (m *CandidateMapper) Insert(newCandidate *entities.Candidate, assessmentId int64) (int64, error) {
	var insertedId int64
	//обращения к БД
	//добавляем кандидата к списку кандидатов
	insertQuery := `INSERT INTO t_candidate 
		(c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_resume, c_addres, c_birth_date, c_education) 
		SELECT nextval('candidate_id'), $1, $2, $3, $4, $5, $6, $7, to_date($8,'YYYY-MM-DD'), $9 
		WHERE NOT EXISTS(SELECT c_id, c_surname, c_name, c_patronymic, c_email, c_phone_number, c_resume, c_addres, c_birth_date, c_education, fk_status FROM t_candidate WHERE c_surname = $10 AND c_name = $11 AND c_patronymic = $12 AND c_birth_date = to_date($13,'YYYY-MM-DD'))
		`
	_, err := m.db.Exec(insertQuery, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.Email, newCandidate.PhoneNumber, newCandidate.Resume, newCandidate.Addres, newCandidate.BirthDate, newCandidate.Education, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.BirthDate)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки кандидата: %v", err)
	}
	//получаем его ID
	row := m.db.QueryRow(`select c_id FROM t_candidate WHERE c_surname = $1 AND c_name = $2 AND c_patronymic = $3 AND c_birth_date = to_date($4,'YYYY-MM-DD')`, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.BirthDate)

	err = row.Scan(&insertedId)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}
	//добавляем кандидата в таблицу связи toc_assessment_candidate
	insertQueryToAssess := `INSERT INTO toc_assessment_candidate 
		(c_id, c_id_assessment, c_id_candidate, c_status_candidate) 
		SELECT nextval('assessment_candidate_id'), $1, $2, 1 
		WHERE NOT EXISTS(SELECT c_id, c_id_assessment, c_id_candidate FROM toc_assessment_candidate WHERE c_id_candidate = $3 AND c_id_assessment = $4)`
	_, err = m.db.Exec(insertQueryToAssess, assessmentId, insertedId, insertedId, assessmentId)

	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки кандидата в ассессмент: %v", err)
	}
	return insertedId, nil
}

//изменение кандидата
func (m *CandidateMapper) Update(newCandidate *entities.Candidate, candidateId int64) (int64, error) {
	//обращение к БД
	insertQuery := `UPDATE t_candidate 
		SET c_surname = $1, c_name = $2, c_patronymic = $3, c_email = $4, c_phone_number = $5, c_resume = $6, c_addres = $7, c_birth_date = to_date($8,'YYYY-MM-DD'), c_education = $9 
		WHERE c_id = $10`
	_, err := m.db.Exec(insertQuery, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.Email, newCandidate.PhoneNumber, newCandidate.Resume, newCandidate.Addres, newCandidate.BirthDate, newCandidate.Education, candidateId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка обновления кандидата: %v", err)
	}
	return candidateId, nil
}

//удаление
func (m *CandidateMapper) Delete(id int64, idAssessment int64) error {
	//обращение к БД
	_, err := m.db.Exec("DELETE FROM toc_assessment_candidate WHERE c_id_candidate = $1 AND c_id_assessment = $2", id, idAssessment)
	fmt.Print("ID Candidate: ", id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("CandidateMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("CandidateMapper::Delete:%v", err)
	}
	return nil
}
