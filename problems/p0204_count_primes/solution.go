package p0204_count_primes

func countPrimes(n int) int {

	slice := make([]bool, n)

	for i := 2; i*i <= n; i++ {
		for j := 2 * i; j < n; j += i {
			slice[j] = true
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !slice[i] {
			count++
		}
	}

	return count

}
