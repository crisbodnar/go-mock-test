package second

//go:generate pegomock generate Person

type Person interface {
	GetName() string
	IsOlderThan(age int) bool
}
