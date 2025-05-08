package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	// Рекурсивная функция для генерации комбинаций
	var generate func(start int, combination []int)
	first := true // Флаг для добавления разделителя ", " только после первой комбинации
	generate = func(start int, combination []int) {
		if len(combination) == n {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false
			printCombination(combination)
			return
		}
		for i := start; i <= 9; i++ {
			generate(i+1, append(combination, i))
		}
	}
	generate(0, []int{})
	z01.PrintRune('\n') // Перевод строки после завершения вывода
}

func printCombination(combination []int) {
	for _, num := range combination {
		z01.PrintRune(rune(num + '0'))
	}
}
