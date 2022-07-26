package teststore

import (
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/store"
)

type Store struct {
	userRepository     *UserRepository
	pharmacyRepository *PharmacyRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
	return s.userRepository
}

func (s *Store) Pharmacy() store.PharmacyRepository {
	if s.pharmacyRepository != nil {
		return s.pharmacyRepository
	}

	s.pharmacyRepository = &PharmacyRepository{
		store:      s,
		pharmacies: make(map[int]*model.Pharmacy),
	}
	return s.pharmacyRepository
}
