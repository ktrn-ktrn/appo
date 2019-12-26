package assessment

import (
	"appo/app/models/entities"
	"appo/app/models/mappers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type AssessmentProvider struct {
	db          *sql.DB
	assessments *mappers.AssessmentMapper
}

func (p *AssessmentProvider) Init() error {
	var err error
	connStr := "user=postgres password=password port=5433 dbname=AssessmentManager sslmode=disable"
	p.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	p.assessments = new(mappers.AssessmentMapper)
	p.assessments.Init(p.db)
	return nil
}

func (p *AssessmentProvider) GetAssessments() ([]*entities.Assessment, error) {
	defer p.db.Close()
	assessments, err := p.assessments.Select()
	if err != nil {
		return nil, err
	}
	return assessments, nil
}

func (p *AssessmentProvider) GetAssessmentById(id int64) (*entities.Assessment, error) {
	defer p.db.Close()
	assessment, err := p.assessments.SelectById(id)
	return assessment, err
}

func (p *AssessmentProvider) GetAssessmentStatus(id int64) ([]*entities.AssessmentStatus, error) {
	defer p.db.Close()
	assessment, err := p.assessments.SelectStatus(id)
	return assessment, err
}

func (p *AssessmentProvider) SetAssessmentStatus(newStatus *entities.AssessmentStatus, statusId int64, assessmentId int64) (int64, error) {
	defer p.db.Close()
	status, err := p.assessments.SetStatus(newStatus, statusId, assessmentId)
	return status, err
}

func (p *AssessmentProvider) PutAssessment(newAssessment *entities.Assessment) (*entities.Assessment, error) {
	defer p.db.Close()
	newAssessment.Status = 1
	id, err := p.assessments.Insert(newAssessment)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("AssessmentProvider::PutAssessment:%v", err)
		}
	}
	createdAsessment, err := p.assessments.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("AssessmentProvider::PutAssessment:%v", err)
	}
	return createdAsessment, nil
}

func (p *AssessmentProvider) PostAssessment(newAssessment *entities.Assessment, assessmentId int64) (*entities.Assessment, error) {
	defer p.db.Close()
	id, err := p.assessments.Update(newAssessment, assessmentId)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("AssessmentProvider::PostAssessment:%v", err)
		}
	}
	updatedAsessment, err := p.assessments.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("AssessmentProvider::PostAssessment:%v", err)
	}
	return updatedAsessment, nil
}

func (p *AssessmentProvider) DeleteAssessment(assessmentId int64) error {
	defer p.db.Close()
	err := p.assessments.Delete(assessmentId)
	if err != nil {
		return fmt.Errorf("AssessmentProvider::DeleteAssessment:%v", err)
	}
	return nil
}
