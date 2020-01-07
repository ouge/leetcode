package leetcode

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	i, j, k := 1, 2, 3 // 点位置
	var values [4]int
	var ret []string
	for {
		// 下标范围判断
		if i > len(s)-3 {
			break
		}
		if j > len(s)-2 {
			i, j, k = i+1, i+2, i+3
			continue
		}
		if k > len(s)-1 {
			j, k = j+1, j+2
			continue
		}

		// 值格式判断
		if i > 1 && s[0] == '0' {
			break
		}
		if j-i > 1 && s[i] == '0' {
			i, j, k = i+1, i+2, i+3
			continue
		}
		if k-j > 1 && s[j] == '0' {
			j, k = j+1, j+2
			continue
		}
		if k < len(s)-1 && s[k] == '0' {
			k = k + 1
			continue
		}

		values[0], _ = strconv.Atoi(s[0:i])
		values[1], _ = strconv.Atoi(s[i:j])
		values[2], _ = strconv.Atoi(s[j:k])
		values[3], _ = strconv.Atoi(s[k:])

		// 值范围判断
		if values[0] > 255 {
			break
		}
		if values[1] > 255 {
			i, j, k = i+1, i+2, i+3
			continue
		}
		if values[2] > 255 {
			j, k = j+1, j+2
			continue
		}
		if values[3] > 255 {
			k = k + 1
			continue
		}
		ret = append(ret, fmt.Sprintf("%s.%s.%s.%s", s[0:i], s[i:j], s[j:k], s[k:]))
		k++
	}
	return ret
}
