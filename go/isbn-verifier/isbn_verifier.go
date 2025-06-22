package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
    sum := 0
    count := 0

    for _, r := range isbn {
        if r == '-' {
            continue
        }
        if count == 9 && (r == 'X' || r == 'x') {
            sum += 10 * (10 - count)
            count++
        } else if unicode.IsDigit(r) {
            sum += int(r - '0') * (10 - count)
            count++
        } else {
            return false
        }
    }

    return count == 10 && sum % 11 == 0
}