package prime

func Nth(n int) (int, bool) {
	var primes []int
	for i := 2; len(primes) < n; i++ {
		j := 0
		for _, prime := range primes {
			if i%prime == 0 {
				break
			}
			j++
		}

		if j == len(primes) {
			primes = append(primes, i)
		}

	}
	if len(primes) == 0 {
		return 0, false
	}
	return primes[len(primes)-1], truee
}
