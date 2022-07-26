package sqlstore

import (
	"time"

	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"
)

type PharmacyRepository struct {
	store *Store
}

func (r *PharmacyRepository) Create(p *model.Pharmacy) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO pharmacies(name,address,geog,add_at,description) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		p.Name,
		p.Address,
		"("+p.Latitude+","+p.Longitute+")",
		time.Now(),
		p.Description,
	).Scan(&p.ID); err != nil {
		return err
	}
	return nil
}
