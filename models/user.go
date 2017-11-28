package models

type User struct {
	ID    string
	Name  string
	Count uint64
}

type Users map[string]*User

type ByCount []User

func (u ByCount) Len() int { return len(u) }

func (u ByCount) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

func (u ByCount) Less(i, j int) bool { return u[i].Count > u[j].Count }
