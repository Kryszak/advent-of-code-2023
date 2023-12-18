package common

func Copy[T any](slice [][]T) [][]T {
	copied := make([][]T, len(slice))
	for i := range slice {
		copied[i] = make([]T, len(slice[i]))
		copy(copied[i], slice[i])
	}
	return copied
}
