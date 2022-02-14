package memory

import (
	"github.com/ingoaf/http-api/adding"
	"github.com/ingoaf/http-api/listing"
)

// Memory storage keeps data in memory
type Storage struct {
	attendees []Attendee
}

// GetAttendee returns an attendee with the specified name
func (m *Storage) GetAttendee(name string) (listing.Attendee, error) {
	var attendee listing.Attendee

	for i := range m.attendees {

		if m.attendees[i].Name == name {
			attendee.Name = m.attendees[i].Name
			attendee.Surname = m.attendees[i].Surname

			return attendee, nil
		}
	}

	return attendee, listing.ErrNotFound
}

// GetAllAttendees returns all attendees
func (m *Storage) GetAllAttendees() []listing.Attendee {
	var attendees []listing.Attendee

	for i := range m.attendees {

		attendee := listing.Attendee{
			Name:    m.attendees[i].Name,
			Surname: m.attendees[i].Surname,
		}

		attendees = append(attendees, attendee)
	}

	return attendees
}

// AddAttendee saves the given attendee in the repository
func (m *Storage) AddAttendee(a adding.Attendee) error {

	newAttendee := Attendee{
		Name:    a.Name,
		Surname: a.Surname,
	}

	m.attendees = append(m.attendees, newAttendee)

	return nil
}
