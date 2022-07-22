package model

type Pharmacy struct {
	Name        string
	PhoneNumber string
	Orders      []Order
	Owner       User
	Remains     float32 // Остаток на счету
	Region      Region
}
