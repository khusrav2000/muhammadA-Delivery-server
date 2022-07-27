package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type OrderedProduct struct {
	Product Product `json:"product"`
	Count   int     `json:"count"`
	Price   float32 `json:"price"`
}

type Order struct {
	ID            int              `json:"id"`
	Products      []OrderedProduct `json:"products"`
	PharmacyId    int              `json:"-"`
	UserName      string           `json:"user_name"`
	UserPhone     string           `json:"user_phone"`
	CreateDate    string           `json:"create_date"`
	DeliveredDate string           `json:"delivered_date"`
	Price         float32          `json:"price"`
}
