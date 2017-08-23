package test_pegomock

//go:generate mockgen first Person

type Person interface {
	GetName() string
	IsOlderThan(age int) bool
}

type Alien struct {}

func (a *Alien) isPersonOlderThan(person Person, age int) bool {
	return person.IsOlderThan(age)
}
