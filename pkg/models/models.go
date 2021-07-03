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
	ErrInvalidID = errors.New("models: id provided was invalid")

	// Error invalid password
	ErrInvalidPassword = errors.New("models: incorrect password provided")

	// Error internal server errors
	ErrInternalServerError = errors.New("models: server error")

	// Error Email has exitst
	ErrEmailHasExist = errors.New("model : email has exists")

	// Error create new user fail
	ErrCreateUserFail = errors.New("nodels : fail to create user")

	// Error create new gallery
	ErrCreateGalleryFail = errors.New("nodels : fail to create gallery")

	// Error Fail to get gallery
	ErrShowGalleryFail = errors.New("models : fail to get gallery")

	// Error fail to update gallery
	ErrorUpdateGalleryFail = errors.New("models : fail to update gallery")

	// Error fail to delete gallery
	ErrorDeleteGalleryFail = errors.New("models : fail to delete gallery")
)

func NewDB() *gorm.DB {
	dsn := "tung:tungbn161296@tcp(127.0.0.1:3360)/gallery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
