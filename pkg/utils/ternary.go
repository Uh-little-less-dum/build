package utils

func Ternary[T any](conditional bool, ifTrue, ifFalse T) T {
	if conditional {
		return ifTrue
	} else {
		return ifFalse
	}
}
