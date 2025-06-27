package sieve

func Sieve(limit int) []int {
    if limit < 2 {
        return []int{}
    }

    sieve := make([]bool, limit+1)
    for i := 2; i * i <= limit; i++ {
        if !sieve[i] {
            for j := i * i; j <= limit; j += i {
                sieve[j] = true
            }
        }
    }

    var primes []int
    for i := 2; i <= limit; i++ {
        if !sieve[i] {
            primes = append(primes, i)
        }
    }

    return primes
}

