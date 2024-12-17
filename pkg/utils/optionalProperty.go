package utils

func GetOptionalProperty[T interface{}](opts []T, defaultVal T) T {
	if len(opts) == 1 {
		return opts[0]
	}
	return defaultVal
}
