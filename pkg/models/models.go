package models

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound = errors.New("models: resource not found")

	// ErrInvalidID is returned when an invalid ID is provided
	// to a method like Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")

	ErrInvalidPassword     = errors.New("models: incorrect password provided")
	ErrInternalServerError = errors.New("models: server error")
	ErrEmailHasExist       = errors.New("model : email has exists")
)

func NewDB() *gorm.DB {
	dsn := "tung:tungbn161296@tcp(127.0.0.1:3360)/gallery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
