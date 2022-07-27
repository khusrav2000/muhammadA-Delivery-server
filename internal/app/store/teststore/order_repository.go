package teststore

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create() error {
	return nil
}
