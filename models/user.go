package models

type User struct {
	ID            string
	Name          string
	Hearts        uint64
	MessageCount  uint64
	HeartsGiven   uint64
	HeartsPerPost float64
	Score         float64
}

type Users map[string]*User

type ByHearts []User

func (u ByHearts) Len() int { return len(u) }

func (u ByHearts) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

func (u ByHearts) Less(i, j int) bool { return u[i].Hearts > u[j].Hearts }

type ByHeartsPerPost []User

func (u ByHeartsPerPost) Len() int { return len(u) }

func (u ByHeartsPerPost) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

func (u ByHeartsPerPost) Less(i, j int) bool { return u[i].HeartsPerPost > u[j].HeartsPerPost }

type ByScore []User

func (u ByScore) Len() int { return len(u) }

func (u ByScore) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

func (u ByScore) Less(i, j int) bool { return u[i].Score > u[j].Score }
