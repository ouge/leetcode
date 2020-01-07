package leetcode

import (
	"bytes"
)

// 简单方法
func reverseWords(s string) string {
	words := bytes.Fields([]byte(s))
	var ret []byte
	for i := len(words) - 1; i >= 0; i-- {
		ret = append(ret, words[i]...)
		ret = append(ret, ' ')
	}
	if len(ret) > 0 {
		ret = ret[:len(ret)-1]
	}
	return string(ret)
}
