package repository

import (
	"database/sql"
	"tugas/domain/model"
)

type UserRepository interface {
	FindByID(id int) (*model.Users, error)
	FindByEmail(email string) (*model.Users, error)
	FindAll() ([]model.Users, error)
	Create(user *model.Users) error
	Update(user *model.Users) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByID(id int) (*model.Users, error) {
	query := "SELECT id, email, username, password, role FROM users WHERE id=$1"
	row := r.db.QueryRow(query, id)

	var user model.Users
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.Users, error) {
	query := "SELECT id, email, username, password, role FROM users WHERE email=$1"
	row := r.db.QueryRow(query, email)

	var user model.Users
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll() ([]model.Users, error) {
	query := "SELECT id, email, username, password, role FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		var u model.Users
		if err := rows.Scan(&u.ID, &u.Email, &u.Username, &u.Password, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) Create(user *model.Users) error {
	query := "INSERT INTO users (email, username, password, role) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, user.Email, user.Username, user.Password, user.Role)
	return err
}

func (r *userRepository) Update(user *model.Users) error {
	query := "UPDATE users SET email=$1, username=$2, password=$3, role=$4 WHERE id=$5"
	_, err := r.db.Exec(query, user.Email, user.Username, user.Password, user.Role, user.ID)
	return err
}

func (r *userRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}
