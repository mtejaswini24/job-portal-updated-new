package database

import (
	"fmt"
	"job-portal-api/config"
	"job-portal-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	cfg := config.GetConfig()
	//dsn := "host=postgres user=postgres password=admin dbname=job-portal-app port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DbConfig.DB_Host, cfg.DbConfig.DB_User, cfg.DbConfig.DB_Pswd, cfg.DbConfig.DB_Name, cfg.DbConfig.DB_Port, cfg.DbConfig.DB_Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Company{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Job{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
