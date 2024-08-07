package tinder

import (
	"github.com/go-playground/validator"
)

const NotFoundStr = "person not found"

var Validate *validator.Validate

type Person struct {
	Name        string `json:"name" validate:"required"`
	Height      int    `json:"height" validate:"required,min=0"`
	Gender      string `json:"gender" validate:"oneof=male female"`
	WantedDates int    `json:"wanted_dates" validate:"required,min=1"`
}

type PersonService interface {
	AddPersonAndMatch(p *Person) ([]*Person, error)
	RemovePerson(name string) error
	QuerySinglePeople(n int) ([]*Person, error)
}

type PersonRepository interface {
	GetAllPeople() []*Person
	GetPersonByName(name string) (*Person, error)
	AddPerson(p *Person) error
	UpdatePerson(p *Person) error
	RemovePerson(name string) error
	GetMatchesForPerson(p *Person) ([]*Person, error)
	UpdateMatchesForPerson(p *Person, matches []*Person) error
}
