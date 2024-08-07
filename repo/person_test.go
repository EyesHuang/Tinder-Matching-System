package repo

import (
	"testing"

	person "tinder"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryRepo(t *testing.T) {
	repo1 := NewMemoryRepo()
	repo2 := NewMemoryRepo()

	assert.Equal(t, repo1, repo2)
}

func TestGetAllPeople(t *testing.T) {
	repo := NewMemoryRepo()
	repo.AddPerson(&person.Person{Name: "Alice"})
	repo.AddPerson(&person.Person{Name: "Bob"})

	people := repo.GetAllPeople()
	assert.Len(t, people, 2)
}

func TestGetPerson_Success(t *testing.T) {
	repo := NewMemoryRepo()
	p := &person.Person{Name: "Charlie"}
	err := repo.AddPerson(p)
	assert.NoError(t, err)

	result, err := repo.GetPersonByName("Charlie")
	assert.NoError(t, err)
	assert.Equal(t, p, result)
}

func TestGetPerson_NotFound(t *testing.T) {
	repo := NewMemoryRepo()
	p := &person.Person{Name: "Charlie"}
	repo.AddPerson(p)
	_, err := repo.GetPersonByName("Test")
	assert.EqualError(t, err, "person not found")
}

func TestAddPerson_Success(t *testing.T) {
	repo := NewMemoryRepo()
	err := repo.AddPerson(&person.Person{Name: "YT"})
	assert.NoError(t, err)
}

func TestAddPerson_AlreadyExist(t *testing.T) {
	repo := NewMemoryRepo()
	repo.AddPerson(&person.Person{Name: "Ling"})

	err := repo.AddPerson(&person.Person{Name: "Ling"})
	assert.EqualError(t, err, "person already exists")
}

func TestUpdatePerson_Success(t *testing.T) {
	repo := NewMemoryRepo()
	repo.AddPerson(&person.Person{Name: "Dave", WantedDates: 5})

	p := &person.Person{Name: "Dave", WantedDates: 10}
	err := repo.UpdatePerson(p)
	assert.NoError(t, err)
	updatedPerson, _ := repo.GetPersonByName("Dave")
	assert.Equal(t, updatedPerson.WantedDates, 10)
}

func TestUpdatePerson_NotFound(t *testing.T) {
	repo := NewMemoryRepo()
	repo.AddPerson(&person.Person{Name: "Dave", WantedDates: 5})

	err := repo.UpdatePerson(&person.Person{Name: "Eve"})
	assert.EqualError(t, err, "person not found")
}

func TestRemovePerson_Success(t *testing.T) {
	repo := NewMemoryRepo()
	repo.AddPerson(&person.Person{Name: "Frank"})
	err := repo.RemovePerson("Frank")
	assert.NoError(t, err)
}

func TestRemovePerson_NotFound(t *testing.T) {
	repo := NewMemoryRepo()
	repo.AddPerson(&person.Person{Name: "Frank"})
	err := repo.RemovePerson("Frank")
	assert.NoError(t, err)

	err = repo.RemovePerson("George")
	assert.EqualError(t, err, "person not found")
}

func TestUpdateMatchesForPerson_Success(t *testing.T) {
	repo := NewMemoryRepo()
	p := &person.Person{Name: "Karen"}
	repo.AddPerson(p)
	err := repo.UpdateMatchesForPerson(p, []*person.Person{{Name: "Leo"}})
	assert.NoError(t, err)
}

func TestUpdateMatchesForPerson_NotFound(t *testing.T) {
	repo := NewMemoryRepo()
	err := repo.UpdateMatchesForPerson(&person.Person{Name: "Test2"}, nil)
	assert.EqualError(t, err, "person not found")
}

func TestGetMatchesForPerson_Success(t *testing.T) {
	repo := NewMemoryRepo()
	p := &person.Person{Name: "Hank"}
	repo.AddPerson(p)
	repo.UpdateMatchesForPerson(p, []*person.Person{{Name: "Ivy"}})

	matches, err := repo.GetMatchesForPerson(p)
	assert.NoError(t, err)
	assert.Equal(t, len(matches), 1)
	assert.Equal(t, "Ivy", matches[0].Name)
}

func TestGetMatchesForPerson_NoMatch(t *testing.T) {
	repo := NewMemoryRepo()

	_, err := repo.GetMatchesForPerson(&person.Person{Name: "Jill"})
	assert.EqualError(t, err, "no matches found for the person")
}
