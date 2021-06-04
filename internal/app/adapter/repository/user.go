package repository

import (
	"errors"

	"github.com/airdb/uic/internal/app/domain"
	"github.com/airdb/uic/internal/app/domain/factory"
	"gorm.io/gorm"
)

// Order is the model of orders
type Order struct {
	ID       string `gorm:"column:order_id;type:uuid"`
	PersonID string `gorm:"type:uuid"`
	// Person   Person  // `gorm:"foreignKey:PersonID;references:ID"`
	// Payment  Payment // `gorm:"foreignkey:OrderID;references:ID"`
}

// User is the repository of domain.User
type User struct {
	ID       uint64
	Username string `gorm:"column:nickname"`
}

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
