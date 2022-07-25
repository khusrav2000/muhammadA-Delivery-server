package sqlstore

import (
	"database/sql"

	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	if err := r.store.db.QueryRow(
		"INSERT INTO users (login, encrypted_password, name, surname) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Login,
		u.EncryptedPassword,
		u.Name,
		u.Surname,
	).Scan(&u.ID); err != nil {
		return err
	}
	var role string
	if err := r.store.db.QueryRow("INSERT INTO user_roles (user_id, role) VALUES ($1, $2) RETURNING role",
		u.ID,
		u.Role,
	).Scan(&role); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(&u.ID,
		&u.Login,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetProfile(u *model.User) error {
	if err := r.store.db.QueryRow(
		"SELECT name, surname FROM users WHERE id = $1",
		u.ID,
	).Scan(
		&u.Name,
		&u.Surname,
	); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password FROM users WHERE login = $1",
		login,
	).Scan(&u.ID,
		&u.Login,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) CheckAccessFor(user *model.User, permission string) (bool, error) {
	rows, err := r.store.db.Query(
		"SELECT role FROM user_roles WHERE user_id = $1",
		user.ID,
	)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return false, err
		}
		var id int = 0
		if err := r.store.db.QueryRow(
			"SELECT id FROM role_permissions WHERE role = $1 AND permission = $2 AND has = TRUE",
			role,
			permission,
		).Scan(&id); err != nil {
			if err != sql.ErrNoRows {
				return false, err
			}
		}
		if id != 0 {
			return true, nil
		}
	}

	return false, nil
}
