package repository

import (
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	maxIdleConn = 2
	maxOpenConn = 5
)

// Connection gets connection of mysql database
func Connection() (db *gorm.DB) {
	dsn := os.Getenv("MAIN_DSN_WRITE")
	if !strings.Contains(dsn, "?") {
		dsn += "?charset=utf8&parseTime=True"
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tab_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,   // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

// User is the repository of domain.User
type User struct {
	ID       uint64
	Username string `gorm:"column:nickname"`
}

type OauthConfig struct {
	gorm.Model
	Provider    string `gorm:"column:provider"`
	ClientID    string `gorm:"column:client_id"`
	CientSecret string `gorm:"column:cient_secret"`
	RedirectURL string `gorm:"column:redirect_url"`
	State       string `gorm:"column:state"`
}
