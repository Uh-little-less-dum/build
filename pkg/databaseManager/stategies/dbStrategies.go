package database_strategies

type DatabaseStrategy int

const (
	FromExisting DatabaseStrategy = iota
	GenerateNew
)
