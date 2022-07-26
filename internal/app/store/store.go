package store

type Store interface {
	User() UserRepository
	Pharmacy() PharmacyRepository
}
