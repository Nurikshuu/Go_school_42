package piscine

func Compact(ptr *[]string) int {
	var result []string

	for _, str := range *ptr {
		if str != "" {
			result = append(result, str)
		}
	}

	*ptr = result

	return len(result)
}
