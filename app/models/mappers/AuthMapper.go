package mappers

import (
	"appo/app/models/entities"
	"database/sql"
	//"fmt"
)

type AuthMapper struct {
	db *sql.DB
}

func (m *AuthMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

func (m *AuthMapper) Login(userName string) (*entities.User, error) {

	return nil, nil
}

func (m *AuthMapper) Logout() (*entities.User, error) {

	return nil, nil
}
