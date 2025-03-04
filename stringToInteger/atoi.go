package stringToInteger

import "strconv"

// MyAtoi преобразует строку в 32-битное целое число со знаком.
// Она обрабатывает ведущие пробелы, знаки плюс и минус, и игнорирует последующие нечисловые символы.
// Если число выходит за пределы 32-битного диапазона, возвращается соответствующая граница.

func MyAtoi(s string) int {
	var str string    // Переменная для хранения числовых символов из строки
	var start = false // Флаг, указывающий, началась ли обработка числовых символов

	// Пробегаем по каждому символу входной строки
	for x := 0; x < len(s); x++ {
		// Если текущий символ не пробел и мы еще не начали обработку чисел, пропускаем его
		if s[x] == ' ' && !start {
			continue
		}
		// Если текущий символ это '+' или '-', и мы еще не начали обработку чисел
		if (s[x] == '-' || s[x] == '+') && !start {
			start = true        // Устанавливаем флаг, что обработка началась
			str += string(s[x]) // Добавляем знак в строку
			continue            // Переходим к следующему символу
		}
		// Если текущий символ это цифра
		if s[x] >= '0' && s[x] <= '9' {
			// Устанавливаем флаг, что обработка началась
			start = true
			str += string(s[x]) // Добавляем цифру в строку

		} else {
			break // Если символ не цифра, прекращаем обработку
		}
	}
	// Если строка пустая или содержит только знак, возвращаем 0
	if str == "" || str == "-" || str == "+" {
		return 0
	}
	// Преобразуем строку в число
	num, err := strconv.Atoi(str)
	if err != nil {
		// Если преобразование не удалось, возвращаем соответствующую границу
		if str[0] == '-' {
			return -2147483648
		}
		return 2147483647
	}
	// Проверяем, не выходит ли число за пределы 32-битного диапазона
	if num < -2147483648 {
		return -2147483648
	}
	if num > 2147483647 {
		return 2147483647
	}
	// Возвращаем результат
	return num
}
