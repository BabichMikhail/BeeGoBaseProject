package utils

func InIntSlice(searchElem int, slice []int) bool {
	for _, elem := range slice {
		if elem == searchElem {
			return true
		}
	}
	return false
}
