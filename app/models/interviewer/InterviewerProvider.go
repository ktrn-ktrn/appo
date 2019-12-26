package interviewer

import (
	"appo/app/models/entities"
	"appo/app/models/mappers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type InterviewerProvider struct {
	db           *sql.DB
	interviewers *mappers.InterviewerMapper
}

func (p *InterviewerProvider) Init() error {
	//подключение к БД
	var err error
	connStr := "user=postgres password=password port=5433 dbname=AssessmentManager sslmode=disable"
	p.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	p.interviewers = new(mappers.InterviewerMapper)
	p.interviewers.Init(p.db)
	return nil
}

//выбрать всех сотрудников которые существуют
func (p *InterviewerProvider) GetAllInterviewers() ([]*entities.Interviewer, error) {
	defer p.db.Close()
	interviewers, err := p.interviewers.SelectAll()
	if err != nil {
		return nil, err
	}
	return interviewers, nil
}

//получить всех сотрудников в выбранном ассессменте
func (p *InterviewerProvider) GetInterviewers(assessmentId int64) ([]*entities.Interviewer, error) {
	defer p.db.Close()
	interviewers, err := p.interviewers.Select(assessmentId)
	if err != nil {
		return nil, err
	}
	return interviewers, nil
}

//получить выбранного сотрудника
func (p *InterviewerProvider) GetInterviewerById(id int64) (*entities.Interviewer, error) {
	defer p.db.Close()
	interviewer, err := p.interviewers.SelectById(id)
	return interviewer, err
}

//добавить сотрудника в ассессмент
func (p *InterviewerProvider) PutInterviewer(newInterviewer *entities.Interviewer, assessmentId int64) (*entities.Interviewer, error) {
	defer p.db.Close()

	//добавляем
	id, err := p.interviewers.Insert(newInterviewer, assessmentId)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("InterviewerProvider::PutInterviewer:%v", err)
		}
	}

	//возвращаем добавленного сотрудника
	createdInterviewer, err := p.interviewers.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("InterviewerProvider::PutInterviewer:%v", err)
	}
	return createdInterviewer, nil
}

//создать сотрудника в словаре
func (p *InterviewerProvider) SetInterviewer(newInterviewer *entities.Interviewer) (*entities.Interviewer, error) {
	defer p.db.Close()
	//создаем сотрудника
	id, err := p.interviewers.InsertInterviewer(newInterviewer)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("InterviewerProvider::SetInterviewer:%v", err)
		}
	}

	//возвращаем созданного сотрудника
	createdInterviewer, err := p.interviewers.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("InterviewerProvider::SetInterviewer:%v", err)
	}
	return createdInterviewer, nil
}

//изменение сотрудника
func (p *InterviewerProvider) PostInterviewer(newInterviewer *entities.Interviewer, InterviewerId int64) (*entities.Interviewer, error) {
	defer p.db.Close()

	//изеняем сотрудника
	id, err := p.interviewers.Update(newInterviewer, InterviewerId)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("InterviewerProvider::PostInterviewer:%v", err)
		}
	}

	//возвращаем изменённого сотрудника
	updatedInterviewer, err := p.interviewers.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("InterviewerProvider::PostInterviewer:%v", err)
	}
	return updatedInterviewer, nil
}

//удаление сотрудника из ассессмента
func (p *InterviewerProvider) DeleteInterviewer(id int64, idAssessment int64) error {
	defer p.db.Close()
	err := p.interviewers.Delete(id, idAssessment)
	if err != nil {
		return fmt.Errorf("InterviewerProvider::DeleteInterviewer:%v", err)
	}
	return err
}

//удаление сотрудника из словаря
func (p *InterviewerProvider) DeleteInterviewerFromD(id int64) error {
	defer p.db.Close()
	err := p.interviewers.DeleteFromD(id)
	if err != nil {
		return fmt.Errorf("InterviewerProvider::DeleteInterviewerFromD:%v", err)
	}
	return err
}
