package piscine

func JumpOver(str string) string {
	var result string
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	if result == "" {
		return "\n"
	}
	return result + "\n"
}
