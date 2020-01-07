package leetcode

func multiply(num1 string, num2 string) string {
	res := make([]byte, len(num1)+len(num2))
	for i := range res {
		res[i] = '0'
	}

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			t := (res[i+j+1] - '0') + (num1[i]-'0')*(num2[j]-'0')
			res[i+j+1] = t%10 + '0'
			res[i+j] += t / 10
		}
	}
	for i := range res {
		if res[i] != '0' {
			return string(res[i:])
		}
	}

	return "0"
}
