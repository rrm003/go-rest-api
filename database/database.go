package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB Database

type Database interface {
	Create(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
}

type GormDatabase struct {
	DB *gorm.DB
}

func (g *GormDatabase) Create(value interface{}) *gorm.DB {
	return g.DB.Create(value)
}

func (g *GormDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	return g.DB.Where(query, args...)
}

func (g *GormDatabase) First(out interface{}, where ...interface{}) *gorm.DB {
	return g.DB.First(out, where...)
}

func (g *GormDatabase) Find(out interface{}, where ...interface{}) *gorm.DB {
	return g.DB.Find(out, where...)
}

func (g *GormDatabase) Save(value interface{}) *gorm.DB {
	return g.DB.Save(value)
}

func (g *GormDatabase) Delete(value interface{}, where ...interface{}) *gorm.DB {
	return g.DB.Delete(value, where...)
}

func InitDatabase() (*GormDatabase, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	DB := &GormDatabase{DB: db}

	return DB, nil
}
