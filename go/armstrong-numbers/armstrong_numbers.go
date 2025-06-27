package armstrong

import "math"

func IsNumber(n int) bool {
    if n < 0 {
        return false
    }

    original := n
    var digits []int

    for n > 0 {
        digits = append(digits, n % 10)
        n /= 10
    }

    power := len(digits)
    sum := 0

    for _, d := range digits {
        sum += int(math.Pow(float64(d), float64(power)))
    }

    return sum == original
}

