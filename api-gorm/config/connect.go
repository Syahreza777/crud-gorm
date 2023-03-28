package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Contact struct {
	Id    uint
	Name  string
	Phone string
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/contact"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Contact{})
	return db, nil
}
