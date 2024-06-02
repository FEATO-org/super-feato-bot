package utility

// 文字配列をユニークにして返す
func StringArrayUnique(arr []string) []string {
	array := make(map[string]struct{})
	unique := []string{}

	for _, ele := range arr {
		array[ele] = struct{}{}
	}
	for arr := range array {
		unique = append(unique, arr)
	}

	return unique
}
