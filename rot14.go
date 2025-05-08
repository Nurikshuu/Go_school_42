package piscine

func Rot14(str string) string {
	runa := []rune(str)
	var result string
	for i := 0; i < len(runa); i++ {
		if runa[i] >= 65 && runa[i] <= 76 {
			runa[i] = runa[i] + 14
		} else if runa[i] >= 109 && runa[i] <= 122 {
			runa[i] = runa[i] - 12
		} else if runa[i] >= 97 && runa[i] <= 108 {
			runa[i] = runa[i] + 14
		} else if runa[i] >= 77 && runa[i] <= 90 {
			runa[i] = runa[i] - 12
		}
		result += string(runa[i])
	}
	return result
}
