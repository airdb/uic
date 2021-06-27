package repository

import (
	"errors"

	"github.com/airdb/uic/internal/app/domain"
	"github.com/airdb/uic/internal/app/domain/factory"
	"gorm.io/gorm"
)

// Get gets order
func (u User) Get() domain.User {
	db := Connection()
	var user User

	result := db.Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	orderFactory := factory.User{}
	return orderFactory.Generate(
		user.ID,
		user.Username,
	)
}

// Update updates order
func (u User) Update(order domain.User) {
	db := Connection()

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("update payments set card_id = 1").Error
		if err != nil {
			return errors.New("rollback")
		}
		return nil // commit
	})
	if err != nil {
		panic(err)
	}
}
