package models

type User struct {
	ID           string
	Name         string
	Hearts       uint64
	MessageCount uint64
}

type Users map[string]*User

type ByHearts []User

func (u ByHearts) Len() int { return len(u) }

func (u ByHearts) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

func (u ByHearts) Less(i, j int) bool { return u[i].Hearts > u[j].Hearts }
