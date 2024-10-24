package db

import (
	"github.com/Mohammedbilal1998/go-postgres/pkg/common/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init (url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Book{})
	return db, nil

}