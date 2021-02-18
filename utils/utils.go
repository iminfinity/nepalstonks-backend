package utils

// PustItemToFirst func
func PustItemToFirst(list []string, newItem string) []string {
	updatedList := []string{}

	updatedList = append(updatedList, newItem)
	for _, item := range list {
		updatedList = append(updatedList, item)
	}

	return updatedList
}
