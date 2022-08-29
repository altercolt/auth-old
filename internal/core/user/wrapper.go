package user

// User for casting model to user
// Mainly used in the repository layer
func (m *Model) User() *User {
	var res User

	if m.ID != nil {
		res.ID = *m.ID
	}

	if m.Email != nil {
		res.Email = *m.Email
	}

	if m.Username != nil {
		res.Username = *m.Username
	}

	if m.Firstname != nil {
		res.Firstname = *m.Firstname
	}

	if m.Lastname != nil {
		res.Lastname = *m.Lastname
	}

	if m.BirthDate != nil {
		res.BirthDate = *m.BirthDate
	}

	return &res
}
