package leetcode

func validUtf8(data []int) bool {
	var i int
	for i < len(data) {
		if data[i]&128 == 0 {
			// 1字节
			i++
		} else if data[i]&224 == 192 {
			// 2字节
			if i+1 >= len(data) || data[i+1]&192 != 128 {
				return false
			}
			i += 2
		} else if data[i]&240 == 224 {
			// 3字节
			if i+2 >= len(data) || data[i+1]&192 != 128 || data[i+2]&192 != 128 {
				return false
			}
			i += 3
		} else if data[i]&248 == 240 {
			// 4字节
			if i+3 >= len(data) || data[i+1]&192 != 128 || data[i+2]&192 != 128 || data[i+3]&192 != 128 {
				return false
			}
			i += 4
		} else {
			return false
		}
	}
	return true
}
