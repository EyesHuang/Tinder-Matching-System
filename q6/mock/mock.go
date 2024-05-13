package mock

import (
	person "tinder"

	"github.com/stretchr/testify/mock"
)

// MockPersonService is a mock type for the PersonService
type MockPersonService struct {
	mock.Mock
}

func (m *MockPersonService) AddPersonAndMatch(p *person.Person) ([]*person.Person, error) {
	args := m.Called(p)
	return args.Get(0).([]*person.Person), args.Error(1)
}

func (m *MockPersonService) RemovePerson(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockPersonService) QuerySinglePeople(n int) ([]*person.Person, error) {
	args := m.Called(n)
	return args.Get(0).([]*person.Person), args.Error(1)
}

// MockPersonService is a mock type for the PersonRepository
type MockPersonRepository struct {
	mock.Mock
}

func (m *MockPersonRepository) GetPerson(name string) (*person.Person, error) {
	args := m.Called(name)
	return args.Get(0).(*person.Person), args.Error(1)
}

func (m *MockPersonRepository) AddPerson(p *person.Person) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockPersonRepository) UpdatePerson(p *person.Person) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockPersonRepository) RemovePerson(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockPersonRepository) GetAllPeople() []*person.Person {
	args := m.Called()
	return args.Get(0).([]*person.Person)
}

func (m *MockPersonRepository) GetMatchesForPerson(p *person.Person) ([]*person.Person, error) {
	args := m.Called(p)
	return args.Get(0).([]*person.Person), args.Error(1)
}

func (m *MockPersonRepository) UpdateMatchesForPerson(p *person.Person, matches []*person.Person) error {
	args := m.Called(p, matches)
	return args.Error(0)
}
