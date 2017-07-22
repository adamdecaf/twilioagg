package phone

import (
	"fmt"
)

type Subject struct {
	Number string
	City, Country, State, Zip string
}
func (s Subject) String() string {
	citystate := s.State
	if s.City != "" {
		citystate = fmt.Sprintf("%s, %s", s.City, s.State)
	}
	if s.Zip != "" {
		citystate += " " + s.Zip
	}
	if s.Country != "" {
		citystate += " " + s.Country
	}
	return fmt.Sprintf("%s from %s", s.Number, citystate)
}
