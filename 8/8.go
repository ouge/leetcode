package leetcode

import "unicode"

const INT_MAX, INT_MIN = 2147483647, -2147483648

func myAtoi(str string) int {
	var i int
	var ret int64
	var flag int64 = 1

	for ; i < len(str) && unicode.IsSpace(rune(str[i])); i++ {
	}
	if i == len(str) || (!unicode.IsDigit(rune(str[i])) && str[i] != '+' && str[i] != '-') {
		return 0
	}
	if str[i] == '+' || str[i] == '-' {
		if str[i] == '-' {
			flag = -1
		}
		i++
	}
	for ; i < len(str) && unicode.IsDigit(rune(str[i])); i++ {
		ret = ret*10 + int64(str[i]-'0')
		if ret*flag > INT_MAX {
			return INT_MAX
		}
		if ret*flag < INT_MIN {
			return INT_MIN
		}
	}
	return int(ret * flag)
}
