package leetcode

var num2char = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := do(digits, nil, nil)
	return res
}

func do(digits string, buffer []byte, res []string) []string {
	if len(digits) == len(buffer) {
		res = append(res, string(buffer))
		return res
	}
	for _, c := range num2char[digits[len(buffer)]] {
		buffer = append(buffer, c)
		res = do(digits, buffer, res)
		buffer = buffer[:len(buffer)-1]
	}
	return res
}
