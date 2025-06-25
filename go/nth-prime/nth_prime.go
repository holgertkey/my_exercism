package prime

import "errors"

func Nth(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("n must be a positive integer")
    }

    count := 0
    for candidate := 2; ; candidate++ {
        if isPrime(candidate) {
            count++
            if count == n {
                return candidate, nil
            }
        }
    }
}

func isPrime(num int) bool {
    if num < 2 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}
