package main

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
	Server *echo
	Store  Store
}

type Store interface {
	createUser(name string, phone int64, email string) error
}
