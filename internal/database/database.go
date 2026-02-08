package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func Connect(dsn string) *gorm.DB {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}


	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Println("failed to connect database")
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("failed to get sql db")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)


	log.Println("connected to database")

	return db

}


func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("failed to get sql db")
		return err
	}

	log.Println("closed database connection")

	return sqlDB.Close()
}