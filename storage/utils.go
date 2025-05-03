package storage

func removeFromSlice[E comparable](slice []E, item E) []E {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice // item not found; return unchanged slice
}
