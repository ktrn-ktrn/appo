package candidate

import (
	"appo/app/models/entities"
	"appo/app/models/mappers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type CandidateProvider struct {
	db         *sql.DB
	candidates *mappers.CandidateMapper
}

func (p *CandidateProvider) Init() error {
	//подключение к БД
	var err error
	connStr := "user=postgres password=password port=5433 dbname=AssessmentManager sslmode=disable"
	p.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	p.candidates = new(mappers.CandidateMapper)
	p.candidates.Init(p.db)
	return nil
}

//получить список всех существующих кандидатов
func (p *CandidateProvider) GetAllCandidates() ([]*entities.Candidate, error) {
	defer p.db.Close()
	candidates, err := p.candidates.SelectAllCandidates()
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

//получем всех кандидатов, участвующих в выбраном ассессменте
func (p *CandidateProvider) GetCandidates(assessmentId int64) ([]*entities.Candidate, error) {
	defer p.db.Close()
	candidates, err := p.candidates.Select(assessmentId)
	if err != nil {
		return nil, fmt.Errorf("CandidateProvider::GetCandidates:%v", err)
	}
	return candidates, nil
}

//получаем выбранного кандидата
func (p *CandidateProvider) GetCandidateById(id int64, assessmentId int64) (*entities.Candidate, error) {
	defer p.db.Close()
	candidate, err := p.candidates.SelectById(id, assessmentId)

	if err != nil {
		return nil, fmt.Errorf("CandidateProvider::GetCandidateById:%v", err)
	}
	return candidate, err
}

//получаем возможные статусы кандидата
func (p *CandidateProvider) GetCandidateStatus(id int64, candidateId int64) ([]*entities.CandidateStatus, error) {
	defer p.db.Close()
	candidate, err := p.candidates.SelectStatus(id, candidateId)
	return candidate, err
}

//задать статус кандидата в выбранном ассессменте
func (p *CandidateProvider) SetCandidateStatus(newStatus *entities.CandidateStatus, statusId int64, assessmentId int64) (int64, error) {
	defer p.db.Close()
	status, err := p.candidates.SetStatus(newStatus, statusId, assessmentId)
	return status, err
}

//добавить кандидата в выбранный ассессмент
func (p *CandidateProvider) PutCandidate(newCandidate *entities.Candidate, assessmentId int64) (*entities.Candidate, error) {
	defer p.db.Close()

	//добавляем кандидата
	id, err := p.candidates.Insert(newCandidate, assessmentId)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("CandidateProvider::PutCandidate:%v", err)
		}
	}

	//возвращаем добавленного
	createdCandidate, err := p.candidates.SelectById(id, assessmentId)
	if err != nil {
		return nil, fmt.Errorf("CandidateProvider::PutCandidate:%v", err)
	}
	return createdCandidate, nil
}

//изменить кандидата
func (p *CandidateProvider) PostCandidate(newCandidate *entities.Candidate, candidateId int64, assessmentId int64) (*entities.Candidate, error) {
	defer p.db.Close()

	//изменяем кадидата
	id, err := p.candidates.Update(newCandidate, candidateId)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("CandidateProvider::PostCandidate:%v", err)
		}
	}

	//возвращаем изменённого кандидата
	updatedCandidate, err := p.candidates.SelectById(id, assessmentId)
	if err != nil {
		return nil, fmt.Errorf("CandidateProvider::PostCandidate:%v", err)
	}
	return updatedCandidate, nil
}

//удалить кандидата из выбраного ассессмента
func (p *CandidateProvider) DeleteCandidate(id int64, idAssessment int64) error {
	defer p.db.Close()
	err := p.candidates.Delete(id, idAssessment)
	if err != nil {
		return fmt.Errorf("CandidateProvider::DeleteCandidate:%v", err)
	}
	return err
}
