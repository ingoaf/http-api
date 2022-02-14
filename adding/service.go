package adding

import (
	"errors"

	"github.com/ingoaf/http-api/listing"
)

// ErrDuplicate is used when a beer already exists.
var ErrDuplicate = errors.New("beer already exists")

// Service provides beer adding operations.
type Service interface {
	AddAttendee(Attendee) error
}

// Repository provides access to attendee repository.
type Repository interface {
	// AddAttendee saves a given attendee to the repository.
	AddAttendee(Attendee) error
	// GetAllAttendees returns all attendees saved in storage.
	GetAllAttendees() []listing.Attendee
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddAttendee persists the given attendee to storage
func (s *service) AddAttendee(a Attendee) error {
	// make sure we don't add any duplicates
	existingBeers := s.r.GetAllAttendees()

	for _, existingAttendee := range existingBeers {
		if a.Name == existingAttendee.Name &&
			a.Surname == existingAttendee.Surname {
			return ErrDuplicate
		}
	}

	s.r.AddAttendee(a)

	return nil
}
