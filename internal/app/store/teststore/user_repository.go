package teststore

import (
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) GetProfile(u *model.User) error {
	_, ok := r.users[u.ID]
	if !ok {
		return store.ErrRecordNotFound
	}

	return nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	for _, u := range r.users {
		if u.Login == login {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) CheckAccessFor(u *model.User, access string) (bool, error) {
	return true, nil
}
