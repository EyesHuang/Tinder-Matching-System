package service

import (
	"errors"
	"testing"

	person "tinder"
	rMock "tinder/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPersonAndMatch_Success(t *testing.T) {
	mockRepo := new(rMock.MockPersonRepository)
	service := NewMatcherService(mockRepo)
	personA := &person.Person{Name: "Alice", Gender: "female", Height: 170, WantedDates: 2}

	mockRepo.On("GetPersonByName", personA.Name).Return((*person.Person)(nil), errors.New(person.NotFoundStr))
	mockRepo.On("AddPerson", personA).Return(nil)
	mockRepo.On("GetAllPeople").Return([]*person.Person{
		{Name: "Bob", Gender: "male", Height: 180, WantedDates: 2},
	})
	mockRepo.On("UpdatePerson", mock.Anything).Return(nil)
	mockRepo.On("RemovePerson", mock.Anything).Return(nil)

	matches, err := service.AddPersonAndMatch(personA)

	assert.NoError(t, err)
	assert.Len(t, matches, 1)
	assert.Equal(t, "Bob", matches[0].Name)
}

func TestAddPersonAndMatch_PersonExists(t *testing.T) {
	mockRepo := new(rMock.MockPersonRepository)
	service := NewMatcherService(mockRepo)
	personA := &person.Person{Name: "Alice", Gender: "female", Height: 170}

	mockRepo.On("GetPersonByName", personA.Name).Return(personA, nil)

	_, err := service.AddPersonAndMatch(personA)

	assert.Error(t, err)
	assert.Equal(t, "person already exists", err.Error())
}

func TestRemovePerson_Success(t *testing.T) {
	mockRepo := new(rMock.MockPersonRepository)
	service := NewMatcherService(mockRepo)
	personName := "Alice"
	alice := &person.Person{Name: "Alice", Gender: "female", Height: 170}
	bob := &person.Person{Name: "Bob", Gender: "male", Height: 180, WantedDates: 2}
	charlie := &person.Person{Name: "Charlie", Gender: "male", Height: 175, WantedDates: 1}
	allPeople := []*person.Person{bob, charlie}

	mockRepo.On("GetPersonByName", personName).Return(alice, nil)

	mockRepo.On("RemovePerson", personName).Return(nil)

	mockRepo.On("GetAllPeople").Return(allPeople)

	mockRepo.On("GetMatchesForPerson", bob).Return([]*person.Person{alice}, nil)
	mockRepo.On("GetMatchesForPerson", charlie).Return([]*person.Person{alice}, nil)

	mockRepo.On("UpdateMatchesForPerson", mock.Anything, []*person.Person{}).Return(nil).Twice()

	err := service.RemovePerson(personName)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRemovePerson_NotFound(t *testing.T) {
	mockRepo := new(rMock.MockPersonRepository)
	service := NewMatcherService(mockRepo)

	mockRepo.On("GetPersonByName", "Alice").Return((*person.Person)(nil), errors.New(person.NotFoundStr))

	err := service.RemovePerson("Alice")

	assert.Error(t, err)
	assert.Equal(t, person.NotFoundStr, err.Error())
}

func TestQuerySinglePeople_Success(t *testing.T) {
	mockRepo := new(rMock.MockPersonRepository)
	service := NewMatcherService(mockRepo)
	people := []*person.Person{
		{Name: "Alice", Gender: "female", Height: 170, WantedDates: 1},
		{Name: "Bob", Gender: "male", Height: 180, WantedDates: 2},
	}
	mockRepo.On("GetAllPeople").Return(people)

	result, err := service.QuerySinglePeople(1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Alice", result[0].Name)
}

func TestQuerySinglePeople_MoreRequestedThanAvailable(t *testing.T) {
	mockRepo := new(rMock.MockPersonRepository)
	service := NewMatcherService(mockRepo)
	people := []*person.Person{
		{Name: "Alice", Gender: "female", Height: 170, WantedDates: 1},
		{Name: "Bob", Gender: "male", Height: 180, WantedDates: 2},
	}

	mockRepo.On("GetAllPeople").Return(people)

	requestedCount := 5

	result, err := service.QuerySinglePeople(requestedCount)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Alice", result[0].Name)
	assert.Equal(t, "Bob", result[1].Name)
}
