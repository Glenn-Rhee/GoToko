package fakers

import (
	"time"

	"github.com/Glenn-Rhee/gotoko/app/models"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID: uuid.New().String(),
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Password: "P4$eJ#8dG@2aR5b!NcL3KpM9dS6fT_gB7hVq*1x",
		RememberToken: "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}