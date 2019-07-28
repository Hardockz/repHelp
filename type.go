package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type Member struct {
	ID   string
	Name string `json:"name"`
	//facebook	string
	Phone int64  `json:"Phone"`
	Email string `json:"Email"`
}

type Rep struct {
	RepName         string
	NumberOfMembers int
	Members         []Member
}

type Service struct {
	Server *echo.Echo
	Store  Store
}

type Store interface {
	createUser(name string, phone int64, email string) error
}

type StoreImpl struct {
	DB *sqlx.DB
}
