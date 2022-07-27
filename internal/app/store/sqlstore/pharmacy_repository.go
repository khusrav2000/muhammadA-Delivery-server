package sqlstore

import (
	"fmt"
	"strings"
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
		LatLngToString(p.Latitude, p.Longitute),
		time.Now(),
		p.Description,
	).Scan(&p.ID); err != nil {
		return err
	}
	return nil
}

func (r *PharmacyRepository) UpdatePharmacy(p *model.Pharmacy, id int) error {
	if err := p.Validate(); err != nil {
		return err
	}
	if err := r.store.db.QueryRow(
		"UPDATE pharmacies SET name=$1, address=$2, geog=$3, description=$4 WHERE id=$5 RETURNING id",
		p.Name,
		p.Address,
		LatLngToString(p.Latitude, p.Longitute),
		p.Description,
		id,
	).Scan(&p.ID); err != nil {
		return err
	}
	return nil
}

func (r *PharmacyRepository) GetPharmacies() (*[]model.Pharmacy, error) {
	rows, err := r.store.db.Query("SELECT id,name,address,geog,description FROM pharmacies")
	if err != nil {
		return nil, err
	}
	pharmacies := make([]model.Pharmacy, 0)
	for rows.Next() {
		p := model.Pharmacy{}
		var geog string
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Address,
			&geog,
			&p.Description,
		); err != nil {
			return nil, err
		}

		p.Latitude, p.Longitute = LatLongFromString(geog)

		pharmacies = append(pharmacies, p)
	}

	return &pharmacies, nil
}

func LatLngToString(lat string, lng string) string {
	return fmt.Sprintf("(%s,%s)", lat, lng)
}

func LatLongFromString(point string) (string, string) {
	s := strings.Split(point, ",")
	lat := strings.Replace(s[0], "(", "", -1)
	lng := strings.Replace(s[1], ")", "", -1)
	return lat, lng

}
