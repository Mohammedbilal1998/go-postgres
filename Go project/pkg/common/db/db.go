package db

import (
	"github.com/Mohammedbilal1998/go-postgres/pkg/common/config"
	"github.com/Mohammedbilal1998/go-postgres/pkg/common/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetClient(config config.Config) (*gorm.DB, error){

	client, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	client.AutoMigrate(&model.Student{})
	client.AutoMigrate(&model.Course{})

	return client, nil
}