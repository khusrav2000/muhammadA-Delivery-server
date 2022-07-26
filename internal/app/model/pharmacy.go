package model

import validation "github.com/go-ozzo/ozzo-validation"

type Pharmacy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Latitude    string `json:"latitute"`
	Longitute   string `json:"longitute"`
	Description string `json:"description"`
}

func (p *Pharmacy) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Address, validation.Required),
		validation.Field(&p.Latitude, validation.Required),
		validation.Field(&p.Longitute, validation.Required),
	)
}
