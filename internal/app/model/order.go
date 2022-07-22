package model

type OrderedProduct struct {
	Product Product
	Count   int
	Price   float32
}

type Order struct {
	ID              int
	OrderedProducts []OrderedProduct
	Pharmacy        Pharmacy
	UserName        string
	UserPhone       string
	OrderDate       string
	DeliveredDate   string
	Price           float32
}
