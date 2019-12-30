package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	return doGenerate(n, 0, 0, nil, nil)
}

func doGenerate(n int, ln, rn int, s []byte, res []string) []string {
	if ln == n {
		for len(s) < 2*n {
			s = append(s, ')')
		}
		return append(res, string(s))
	}
	res = doGenerate(n, ln+1, rn, append(s, '('), res)
	if ln > rn {
		res = doGenerate(n, ln, rn+1, append(s, ')'), res)
	}
	return res
}
