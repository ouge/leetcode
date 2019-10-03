package leetcode

func simplifyPath(path string) string {
	var ret []byte
	path = "/" + path + "/"
	var dotCnt int

	for _, c := range []byte(path) {
		switch c {
		case '/':
			if len(ret) == 0 {
				ret = append(ret, '/')
				continue
			}
			if ret[len(ret)-1] == '/' {
				// 忽略连续/
				continue
			}

			if dotCnt == 0 {
				ret = append(ret, '/')
			} else if dotCnt > 2 {
				// 当正常filename处理
				ret = append(ret, '/')
				dotCnt = 0
			} else if dotCnt == 1 {
				// 删除末尾 .
				ret = ret[:len(ret)-1]
				dotCnt = 0
			} else if dotCnt == 2 {
				// 删除末尾 /..
				ret = ret[:len(ret)-3]
				// 删除最后一个filename
				for i := len(ret) - 1; i >= 0 && ret[i] != '/'; i-- {
					ret = ret[:len(ret)-1]
				}
				if len(ret) == 0 {
					ret = append(ret, '/')
				}
			} else {
				return ""
			}
			dotCnt = 0
		case '.':
			dotCnt++
			ret = append(ret, '.')
		default:
			ret = append(ret, c)
			dotCnt = 3
		}
	}

	// 删除最后一个/
	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}

	return string(ret)
}
