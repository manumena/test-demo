package students

import "strings"

// Public model returned by the Go service
type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Class     string `json:"class"`
	Section   string `json:"section"`
}

// Internal DTOs matching the Node.js backend response

type studentDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Class *struct {
		Name string `json:"name"`
	} `json:"class"`
	Section *struct {
		Name string `json:"name"`
	} `json:"section"`
}

// Mapper from backend DTO to public model
func toStudent(dto studentDTO) *Student {
	first := dto.Name
	last := ""

	if parts := strings.SplitN(dto.Name, " ", 2); len(parts) == 2 {
		first = parts[0]
		last = parts[1]
	}

	s := &Student{
		ID:        dto.ID,
		FirstName: first,
		LastName:  last,
		Email:     dto.Email,
	}

	if dto.Class != nil {
		s.Class = dto.Class.Name
	}

	if dto.Section != nil {
		s.Section = dto.Section.Name
	}

	return s
}
