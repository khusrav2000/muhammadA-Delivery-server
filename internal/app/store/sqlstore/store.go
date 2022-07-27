package sqlstore

import (
	"database/sql"

	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db                 *sql.DB
	userRepository     *UserRepository
	pharmacyRepository *PharmacyRepository
	orderRepository    *OrderRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

func (s *Store) Pharmacy() store.PharmacyRepository {
	if s.pharmacyRepository != nil {
		return s.pharmacyRepository
	}
	s.pharmacyRepository = &PharmacyRepository{
		store: s,
	}
	return s.pharmacyRepository
}

func (s *Store) Order() store.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}
	s.orderRepository = &OrderRepository{
		store: s,
	}
	return s.orderRepository
}
