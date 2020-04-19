package utils

func ContainsAnyStr(arr []string, elem ...string) int {
	for _, str := range elem {
		index := ContainsStr(arr, str)
		if index != -1 {
			return index
		}
	}
	return -1
}

func ContainsStr(arr []string, elem string) int {
	for i, str := range arr {
		if str == elem {
			return i
		}
	}
	return -1
}
