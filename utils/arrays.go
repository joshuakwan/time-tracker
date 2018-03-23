package utils

import "reflect"

func RemoveElement(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func Contains(element interface{}, slice interface{}) int {
	index := -1
	true_slice := reflect.ValueOf(slice)

	if true_slice.Kind() == reflect.Slice {
		for i := 0; i < true_slice.Len(); i++ {
			if true_slice.Index(i).Interface() == element {
				index = i
			}
		}
	}

	return index
}
