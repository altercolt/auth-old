package user

import (
	"net/mail"
	"time"
)

// Validate called in handlers layer
func (n New) Validate() map[string]string {

	res := make(map[string]string)
	var resBool bool

	if _, err := mail.ParseAddress(n.Email); err != nil {
		res["email"] = err.Error()
		resBool = true
	}

	if n.Username == "" {
		res["username"] = "empty username"
		resBool = true
	} else if n.Username != "" && len(n.Username) < 4 {
		res["username"] = "length of your username must be more than 3 characters"
		resBool = true
	}

	if n.Firstname == "" {
		res["firstname"] = "empty firstname"
		resBool = true
	}

	if n.Lastname == "" {
		res["lastname"] = "empty lastname"
		resBool = true
	}

	if len(n.Password) < 8 {
		res["password"] = "invalid length, password length must be more than 8 characters"
		resBool = true
	}

	if n.BirthDate == (time.Time{}) {
		res["birthdate"] = "empty birth date"
		resBool = true
	}

	if !resBool {
		res = nil
	}

	return res
}

// Validate
//TODO: WRITE VALIDATION FOR UPDATE
func (u Update) Validate() map[string]string {
	res := make(map[string]string)
	return res
}
