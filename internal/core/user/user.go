package user

import "time"

type User struct {
	ID        int       `json:"id"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	BirthDate time.Time `json:"birthDate"`
	PassHash  string    `json:"-"`
}

type New struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	BirthDate time.Time `json:"birthDate"`
	Password  string    `json:"password"`
}

type Update struct {
	Email     *string    `json:"email,omitempty"`
	Username  *string    `json:"username,omitempty"`
	Firstname *string    `json:"firstname,omitempty"`
	Lastname  *string    `json:"lastname,omitempty"`
	BirthDate *time.Time `json:"birthDate,omitempty"`
	Password  *string    `json:"password,omitempty"`
}

type Model struct {
	ID        *int
	Email     *string
	Role      *string
	Username  *string
	Firstname *string
	Lastname  *string
	BirthDate *time.Time
	PassHash  *string
}
