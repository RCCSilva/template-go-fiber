package domain

import "rccsilva.com/template-go/database"

type MemberRepository struct {
	db *database.Database
}

func (r MemberRepository) Create(user *user) (*user, error) {
	query := `
	INSERT INTO users (email, username, password) 
	VALUES ($1, $2, $3)
	RETURNING id;`
	row := r.db.DB.QueryRow(query, user.Email, user.Username, user.Password)

	var err = row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r MemberRepository) Get(id int) (*user, error) {
	query := `
	SELECT email, username
	FROM users
	WHERE id = $1;`

	user := &user{ID: id}

	row := r.db.DB.QueryRow(query, id)

	var err = row.Scan(&user.Email, &user.Username)

	if err != nil {
		return nil, err
	}

	return user, nil
}
