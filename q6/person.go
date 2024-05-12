package tinder

import "time"

type Person struct {
	Name        string      `json:"name"`
	Height      int         `json:"height"`
	Gender      Gender      `json:"gender"`
	WantedDates []time.Time `json:"wanted_dates"`
}

type Gender int

const (
	Male Gender = iota
	Female
)

func (g Gender) String() string {
	names := [...]string{
		"male",
		"female",
	}

	if int(g) < len(names) {
		return names[g]
	}

	return "unknown"
}

type PersonService interface {
	AddPersonAndMatch(p *Person) ([]*Person, error)
	RemovePerson(name string) error
	QuerySinglePeople(n int) ([]*Person, error)
}

type PersonRepository interface {
	GetAllPeople() ([]*Person, error)
	GetPerson(name string) (*Person, error)
	AddPerson(p *Person) error
	UpdatePerson(p *Person) error
	RemovePerson(name string) error
	GetMatchesForPerson(p *Person) ([]*Person, error)
	UpdateMatchesForPerson(p *Person, matches []*Person) error
}
