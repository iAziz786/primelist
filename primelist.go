package primelist

import "math"

func primelist(size int) []int {
	primes := make([]bool, size)

	for i := range primes {
		primes[i] = true
	}

	sqrt := int(math.Sqrt(float64(size)))

	for i := 2; i <= sqrt; i++ {
		for j := 2; j*i < size; j++ {
			if primes[i*j] {
				primes[i*j] = false
			}
		}
	}

	var arrayOfPrimes []int

	for i := 2; i < size; i++ {
		if primes[i] {
			arrayOfPrimes = append(arrayOfPrimes, i)
		}
	}
	
	return arrayOfPrimes
}