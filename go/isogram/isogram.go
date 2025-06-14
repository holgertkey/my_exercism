package isogram

import (
    "strings"
    "unicode"
)

func IsIsogram(word string) bool {
    lowerWord := strings.ToLower(word)
    for _, char := range lowerWord {
        if !unicode.IsLetter(char) {
            continue
        }
        tmp := strings.ReplaceAll(lowerWord, string(char), "")
        if len(tmp) < len(lowerWord)-1 {
            return false
        }
    }
    return true
}
