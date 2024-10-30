package bootstrap

func Coalesce[T comparable](values ...T) T {
	var empty T
	for _, value := range values {
		if value != empty {
			return value
		}
	}
	return empty
}
