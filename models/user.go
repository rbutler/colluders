package models

type User struct {
	ID    string
	Name  string
	Count uint64
}

type Users map[string]*User
