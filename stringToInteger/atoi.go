package stringToInteger

import "strconv"

func MyAtoi(s string) int {
	var str string
	var start = false

	for x := 0; x < len(s); x++ {
		if s[x] == ' ' && !start {
			continue
		}

		if (s[x] == '-' || s[x] == '+') && !start {
			start = true
			str += string(s[x])
			continue
		}

		if s[x] >= '0' && s[x] <= '9' {
			start = true
			str += string(s[x])
		} else {
			break
		}
	}

	if str == "" || str == "-" || str == "+" {
		return 0
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		if str[0] == '-' {
			return -2147483648
		}
		return 2147483647
	}

	if num < -2147483648 {
		return -2147483648
	}
	if num > 2147483647 {
		return 2147483647
	}
	return num
}
