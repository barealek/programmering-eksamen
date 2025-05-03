package storage

func removeFromSlice(slice []*Entry, item *Entry) []*Entry {
	for i, v := range slice {
		if v.ID == item.ID {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice // item not found; return unchanged slice
}
