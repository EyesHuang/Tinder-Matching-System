package service

import (
	"errors"
	"fmt"
	"sort"

	person "tinder"
)

type MatcherService struct {
	repo person.PersonRepository
}

func NewMatcherService(repo person.PersonRepository) *MatcherService {
	return &MatcherService{repo: repo}
}

func (ms *MatcherService) AddPersonAndMatch(p *person.Person) ([]*person.Person, error) {
	// Check if the person already exists
	_, err := ms.repo.GetPerson(p.Name)
	if err == nil {
		return nil, errors.New("person already exists")
	}

	// Add the person to the repository
	if err = ms.repo.AddPerson(p); err != nil {
		return nil, err
	}

	// Find potential matches
	potentialMatches := ms.findPotentialMatches(p)
	matches := make([]*person.Person, 0, len(potentialMatches))

	for _, match := range potentialMatches {
		// Update the wanted dates
		match.WantedDates--
		p.WantedDates--

		// Update the person in the repository
		err = ms.repo.UpdatePerson(match)
		if err != nil {
			return nil, err
		}

		err = ms.repo.UpdatePerson(p)
		if err != nil {
			return nil, err
		}

		// Remove the person from the matching system if their wanted dates become empty
		if match.WantedDates == 0 {
			err = ms.repo.RemovePerson(match.Name)
			if err != nil {
				return nil, err
			}
		}

		if p.WantedDates == 0 {
			err = ms.repo.RemovePerson(p.Name)
			if err != nil {
				return nil, err
			}
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func (ms *MatcherService) RemovePerson(name string) error {
	// Get the person from the repository
	p, err := ms.repo.GetPerson(name)
	if err != nil {
		return err
	}

	if p == nil {
		return fmt.Errorf(person.NotFoundStr)
	}

	// Remove the person from the repository
	err = ms.repo.RemovePerson(name)
	if err != nil {
		return err
	}

	// Update the matches for other people
	allPeople, _ := ms.repo.GetAllPeople()
	for _, other := range allPeople {
		matches, _ := ms.repo.GetMatchesForPerson(other)
		newMatches := make([]*person.Person, 0, len(matches))
		for _, match := range matches {
			if match.Name != p.Name {
				newMatches = append(newMatches, match)
			}
		}
		err = ms.repo.UpdateMatchesForPerson(other, newMatches)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ms *MatcherService) QuerySinglePeople(n int) ([]*person.Person, error) {
	// Get all people from the repository
	allPeople, err := ms.repo.GetAllPeople()
	if err != nil {
		return nil, err
	}

	// Sort the people based on the length of their WantedDates slice (ascending order)
	sort.Slice(allPeople, func(i, j int) bool {
		return allPeople[i].WantedDates < allPeople[j].WantedDates
	})

	// Return the most N possible matched single people
	if n > len(allPeople) {
		return allPeople, nil
	}
	return allPeople[:n], nil
}

func (ms *MatcherService) findPotentialMatches(p *person.Person) []*person.Person {
	allPeople, _ := ms.repo.GetAllPeople()
	potentialMatches := make([]*person.Person, 0)
	for _, other := range allPeople {
		if p.Gender != other.Gender && meetsHeightRequirement(p, other) {
			potentialMatches = append(potentialMatches, other)
		}
	}
	return potentialMatches
}

func meetsHeightRequirement(p1, p2 *person.Person) bool {
	if p1.Gender == "male" {
		return p1.Height > p2.Height
	}
	return p1.Height < p2.Height
}
