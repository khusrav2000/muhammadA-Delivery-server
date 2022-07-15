package store_test

import (
	"log"
	"testing"

	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	log.Println("DEFER FUNC")
	u, err := s.User().Create(&model.User{
		Email: "user@mail.ru",
	})
	log.Println(t)
	log.Println(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
