package entities

type Candidate struct {
	ID          int64
	Surname     string
	Name        string
	Patronymic  string
	Email       string
	PhoneNumber string
	Resume      string
	Addres      string
	Education   string
	BirthDate   string
	Status      int64
	StatusName  string
}
