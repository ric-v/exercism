package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input should be greater than or equal to 0")
	}

	var count int
	for n != 1 {
		if n%2 == 0 {
			n = (n / 2)
		} else if n%2 == 1 {
			n = (3 * n) + 1
		}
		count++
	}
	return count, nil
}
