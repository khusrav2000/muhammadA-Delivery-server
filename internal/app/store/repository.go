package store

import "github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
}
