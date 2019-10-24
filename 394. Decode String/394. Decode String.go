package main

type node struct {
	n   int
	res []byte
}

func decodeString(s string) string {
	stack := []node{}
	b := []byte(s)
	var n int
	var res []byte
	for i := range b {
		switch {
		case '0' <= b[i] && b[i] <= '9':
			n = n*10 + int(b[i]-'0')
		case b[i] == '[':
			stack = append(stack, node{n, res})
			res = nil
			n = 0
		case b[i] == ']':
			for stack[len(stack)-1].n > 0 {
				stack[len(stack)-1].res = append(stack[len(stack)-1].res, res...)
				stack[len(stack)-1].n--
			}
			res = stack[len(stack)-1].res
			stack = stack[:len(stack)-1]
		default:
			res = append(res, b[i])
		}

	}
	return string(res)
}

func main() {

}
