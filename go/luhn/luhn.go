package luhn

import (
    "unicode"
)

func Valid(id string) bool {
    var digits []int
    for _, r := range id {
        if unicode.IsSpace(r) {
            continue
        }
        if !unicode.IsDigit(r) {
            return false
        }
        digits = append(digits, int(r-'0'))
    }

    if len(digits) < 2 {
        return false
    }

    sum := 0
    double := false
    for i := len(digits) - 1; i >= 0; i-- {
        n := digits[i]
        if double {
            n *= 2
            if n > 9 {
                n -= 9
            }
        }
        sum += n
        double = !double 
    }

    return sum%10 == 0
}
