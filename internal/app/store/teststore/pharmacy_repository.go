package teststore

import "github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"

type PharmacyRepository struct {
	store      *Store
	pharmacies map[int]*model.Pharmacy
}

func (r *PharmacyRepository) Create(p *model.Pharmacy) error {
	p.ID = len(r.pharmacies) + 1
	r.pharmacies[p.ID] = p

	return nil
}
