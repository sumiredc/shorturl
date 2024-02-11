package num

import "errors"

// 順列
func Permutation(n int, r int) (int, error) {
	if n < r {
		return 0, errors.New("Please specify a number larger than n for r.")
	}

	p := 1
	cnt := 0
	for r > cnt {
		p *= (n - cnt)
		cnt++
	}
	return p, nil
}
