package repo

import (
	"errors"
	"sync"

	person "tinder"
)

type MemoryRepo struct {
	sync.Mutex
	people  map[string]*person.Person
	matches map[*person.Person][]*person.Person
}

var (
	instance *MemoryRepo
	once     sync.Once
)

func NewMemoryRepo() *MemoryRepo {
	once.Do(func() {
		instance = &MemoryRepo{
			people:  make(map[string]*person.Person),
			matches: make(map[*person.Person][]*person.Person),
		}
	})
	return instance
}

func (r *MemoryRepo) GetAllPeople() []*person.Person {
	r.Lock()
	defer r.Unlock()

	people := make([]*person.Person, 0, len(r.people))
	for _, p := range r.people {
		people = append(people, p)
	}
	return people
}

func (r *MemoryRepo) GetPersonByName(name string) (*person.Person, error) {
	r.Lock()
	defer r.Unlock()

	// strings.ToLower(name)
	p, ok := r.people[name]
	if !ok {
		return nil, errors.New("person not found")
	}
	return p, nil
}

func (r *MemoryRepo) AddPerson(p *person.Person) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.people[p.Name]; exists {
		return errors.New("person already exists")
	}
	r.people[p.Name] = p
	return nil
}

func (r *MemoryRepo) UpdatePerson(p *person.Person) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.people[p.Name]; !exists {
		return errors.New("person not found")
	}
	r.people[p.Name] = p
	return nil
}

func (r *MemoryRepo) RemovePerson(name string) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.people[name]; !exists {
		return errors.New("person not found")
	}
	delete(r.people, name)
	delete(r.matches, r.people[name])
	return nil
}

func (r *MemoryRepo) GetMatchesForPerson(p *person.Person) ([]*person.Person, error) {
	r.Lock()
	defer r.Unlock()

	matches, ok := r.matches[p]
	if !ok {
		return nil, errors.New("no matches found for the person")
	}
	return matches, nil
}

func (r *MemoryRepo) UpdateMatchesForPerson(p *person.Person, matches []*person.Person) error {
	r.Lock()
	defer r.Unlock()

	_, exists := r.people[p.Name]
	if !exists {
		return errors.New("person not found")
	}
	r.matches[p] = matches
	return nil
}
