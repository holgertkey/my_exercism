package atbash

import (
    "strings"
    "unicode"
)

func Atbash(s string) string {
    var result strings.Builder
    count := 0

    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            if count > 0 && count % 5 == 0 {
                result.WriteRune(' ')
            }

            if unicode.IsLetter(r) {
                lower := unicode.ToLower(r)
                mirrored := 'a' + ('z' - lower)
                result.WriteRune(mirrored)
            } else {
                result.WriteRune(r)
            }

            count++
        }
    }

    return result.String()
}

