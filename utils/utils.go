package utils

import "sync"

// PustItemToFirst func
func PustItemToFirst(list []string, newItem string) []string {
	updatedList := []string{}

	updatedList = append(updatedList, newItem)
	for _, item := range list {
		updatedList = append(updatedList, item)
	}

	return updatedList
}

// Contains func
func Contains(list []string, search string) bool {
	for _, item := range list {
		if search == item {
			return true
		}
	}
	return false
}

// Reverse func
func Reverse(wg *sync.WaitGroup, list []string) {
	defer wg.Done()
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
