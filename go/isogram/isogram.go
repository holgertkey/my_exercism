package isogram

import (
    "strings"
    "unicode"
)

func IsIsogram(word string) bool {
    lowerWord := strings.ToLower(word)
    for i, char := range lowerWord {
        if unicode.IsLetter(char) && strings.ContainsRune(lowerWord[i + 1:], char) {
            return false
        }
    }
    return true
}
