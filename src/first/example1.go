package first

//go:generate mockgen first Person

type Person interface {
	GetName() string
	IsOlderThan(age int) bool
}
