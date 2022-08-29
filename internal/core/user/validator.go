package user

import (
	"net/mail"
	"time"
)

// Validate called in handlers layer
func (n *New) Validate() (map[string]string, bool) {

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

	return res, !resBool
}

func (u *Update) Validate() (map[string]string, bool) {
	res := make(map[string]string)

	var resBool bool

	if u.Email != nil {
		if _, err := mail.ParseAddress(*u.Email); err != nil {
			res["email"] = err.Error()
			resBool = true
		}
	}

	if u.Username != nil {
		if *u.Username == "" {
			res["username"] = "empty username"
			resBool = true
		} else if *u.Username != "" && len(*u.Username) < 4 {
			res["username"] = "length of your username must be more than 3 characters"
			resBool = true
		}
	}

	if u.Firstname != nil {
		if *u.Firstname == "" {
			res["firstname"] = "empty firstname"
			resBool = true
		}
	}

	if u.Lastname != nil {
		if *u.Lastname == "" {
			res["lastname"] = "empty lastname"
			resBool = true
		}
	}

	if u.Password != nil {
		if len(*u.Password) < 8 {
			res["password"] = "invalid length, password length must be more than 8 characters"
			resBool = true
		}
	}

	if u.BirthDate != nil {
		if *u.BirthDate == (time.Time{}) {
			res["birthdate"] = "empty birth date"
			resBool = true
		}
	}

	if !resBool {
		res = nil
	}

	return res, resBool
}
