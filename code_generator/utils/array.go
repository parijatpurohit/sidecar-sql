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

func GetUnique(arr []string) []string {
	data := map[string]struct{}{}
	for _, s := range arr {
		data[s] = struct{}{}
	}
	var filtered []string
	for s, _ := range data {
		filtered = append(filtered, s)
	}
	return filtered
}
