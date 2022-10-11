package postgres

import (
	"context"
	"errors"
	"github.com/altercolt/auth/internal/core/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

// userRepository
// implements user.Repository interface
type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) user.Repository {
	return userRepository{
		db: conn,
	}
}

// Create
// used for creating new user record in users table
func (u userRepository) Create(ctx context.Context, m *user.Model) error {
	query := `INSERT INTO users (email, username, firstname, lastname, birth_date, salt, passhash, role_id) 
			VALUES($1, $2, $3, $4, $5, $6, $7, $8);
`
	//TODO: CHANGE ROLE_USER to role.User
	_, err := u.db.Exec(ctx, query,
		m.Email, m.Username, m.Firstname, m.Lastname, m.BirthDate,
		m.Salt, m.PassHash, 2)

	if err != nil {
		return err
	}

	return nil
}

// Update
// used for updating user record in users table
// TODO: (IF NOT NULL UPDATE, m.Username is NOT NULL UPDATE, etc)
func (u userRepository) Update(ctx context.Context, m *user.Model) error {
	query := `UPDATE users
			  SET email = COALESCE($2, email),
			      username = COALESCE($3, username),
			      firstname = COALESCE($4, firstname),
			      lastname = COALESCE($5, lastname),
			      birth_date = COALESCE($6, birth_date),
			      salt = COALESCE($7, salt),
			      passhash = COALESCE($8, passhash)
			WHERE id = $1;
			      `
	tx, err := u.db.Begin(ctx)
	res, err := u.db.Exec(ctx, query,
		m.ID,
		m.Email, m.Username, m.Firstname, m.Lastname, m.BirthDate,
		m.Salt, m.PassHash)

	if res.RowsAffected() > 1 {
		err = tx.Rollback(ctx)
		return err
	}

	return err
}

// Fetch
// used for getting multiple users from users table
func (u userRepository) Fetch(ctx context.Context, f user.Filter) ([]user.User, error) {
	query := `SELECT users.id, 
				roles.name, users.email, 
				users.username, users.firstname, users.lastname, users.birth_date, 
				users.salt, users.passhash
				FROM users 
				INNER JOIN roles 
				    ON users.role_id = roles.id
				WHERE ($1::int[] IS NULL OR users.id = ANY($1))
         		AND ($2::varchar[] IS NULL OR users.email = ANY($2))
         		AND ($3::varchar[] IS NULL OR users.username = ANY($3))
         		AND ($4::varchar[] IS NULL OR users.firstname = ANY($4))
         		AND ($5::varchar[] IS NULL OR users.lastname = ANY($5))
         		AND ($6::date[] IS NULL OR users.birth_date = ANY($6));`

	rows, err := u.db.Query(ctx, query, f.ID, f.Email, f.Username,
		f.Firstname, f.Lastname, f.BirthDate)
	if err != nil {
		return nil, err
	}

	var res []user.User

	for rows.Next() {
		var usr user.User
		if err = rows.Scan(&usr.ID, &usr.Role, &usr.Email, &usr.Username, &usr.Firstname, &usr.Lastname, &usr.BirthDate, &usr.Salt, &usr.PassHash); err != nil {
			return nil, err
		}

		res = append(res, usr)
	}

	return res, nil
}

// FetchOne
// used for getting a single user from users table
func (u userRepository) FetchOne(ctx context.Context, f user.Filter) (user.User, error) {
	query := `SELECT users.id, 
				roles.name, users.email, 
				users.username, users.firstname, users.lastname, users.birth_date, 
				users.salt, users.passhash
				FROM users 
				INNER JOIN roles 
				    ON users.role_id = roles.id
				WHERE ($1::int[] IS NULL OR users.id = ANY($1))
         		AND ($2::varchar[] IS NULL OR users.email = ANY($2))
         		AND ($3::varchar[] IS NULL OR users.username = ANY($3))
         		AND ($4::varchar[] IS NULL OR users.firstname = ANY($4))
         		AND ($5::varchar[] IS NULL OR users.lastname = ANY($5))
         		AND ($6::date[] IS NULL OR users.birth_date = ANY($6));`

	var usr user.User

	if err := validateFilter(f); err != nil {
		return usr, err
	}

	row := u.db.QueryRow(ctx, query, f.ID, f.Email, f.Username, f.Firstname, f.Lastname, f.BirthDate)

	if err := row.Scan(&usr.ID, &usr.Role, &usr.Email, &usr.Username, &usr.Firstname, &usr.Lastname, &usr.BirthDate, &usr.Salt, &usr.PassHash); err != nil {
		return usr, err
	}

	return usr, nil
}

// Delete
// used for deleting a single user
func (u userRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	tx, err := u.db.Begin(ctx)
	if err != nil {
		return err
	}

	res, err := u.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if res.RowsAffected() > 1 {
		err = tx.Rollback(ctx)
		if err != nil {
			return err
		}

		return err
	}

	return nil
}

func validateFilter(f user.Filter) error {

	if len(f.ID) > 1 || len(f.Email) > 1 || len(f.Username) > 1 || len(f.Firstname) > 1 ||
		len(f.Lastname) > 1 || len(f.BirthDate) > 1 {
		return errors.New("length of filters in FetchOne() cannot be more than 1 ")
	}

	return nil
}
