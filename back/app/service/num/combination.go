package num

import "errors"

// 順列
func Permutation(n int, r int) (int, error) {
	if n < r {
		return 0, errors.New("")
	}

	p := 1
	for n > r {
		p *= n
		n--
	}
	return p, nil
}
