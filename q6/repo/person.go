package db

import (
	"errors"

	person "tinder"
)

type MemoryRepo struct {
	people  map[string]*person.Person
	matches map[*person.Person][]*person.Person
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		people:  make(map[string]*person.Person),
		matches: make(map[*person.Person][]*person.Person),
	}
}

func (r *MemoryRepo) GetAllPeople() ([]*person.Person, error) {
	people := make([]*person.Person, 0, len(r.people))
	for _, p := range r.people {
		people = append(people, p)
	}
	return people, nil
}

func (r *MemoryRepo) GetPerson(name string) (*person.Person, error) {
	p, ok := r.people[name]
	if !ok {
		return nil, errors.New("person not found")
	}
	return p, nil
}

func (r *MemoryRepo) AddPerson(p *person.Person) error {
	if _, exists := r.people[p.Name]; exists {
		return errors.New("person already exists")
	}
	r.people[p.Name] = p
	return nil
}

func (r *MemoryRepo) UpdatePerson(p *person.Person) error {
	if _, exists := r.people[p.Name]; !exists {
		return errors.New("person not found")
	}
	r.people[p.Name] = p
	return nil
}

func (r *MemoryRepo) RemovePerson(name string) error {
	if _, exists := r.people[name]; !exists {
		return errors.New("person not found")
	}
	delete(r.people, name)
	delete(r.matches, r.people[name])
	return nil
}

func (r *MemoryRepo) GetMatchesForPerson(p *person.Person) ([]*person.Person, error) {
	matches, ok := r.matches[p]
	if !ok {
		return nil, nil // No matches found for the person
	}
	return matches, nil
}

func (r *MemoryRepo) UpdateMatchesForPerson(p *person.Person, matches []*person.Person) error {
	_, exists := r.people[p.Name]
	if !exists {
		return errors.New("person not found")
	}
	r.matches[p] = matches
	return nil
}
