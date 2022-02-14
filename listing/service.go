package listing

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("attendee not found")

// Repository provides access to the beer and review storage.
type Repository interface {
	// GetAttendee returns the attendee with given name.
	GetAttendee(string) (Attendee, error)
	// GetAllAttendees returns all attendees saved in storage.
	GetAllAttendees() []Attendee
}

// Service provides attendee listing operations.
type Service interface {
	GetAttendee(string) (Attendee, error)
	GetAttendees() []Attendee
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetAttendees returns all attendies
func (s *service) GetAttendees() []Attendee {
	return s.r.GetAllAttendees()
}

// GetAttendee returns an attendee
func (s *service) GetAttendee(name string) (Attendee, error) {
	return s.r.GetAttendee(name)
}
