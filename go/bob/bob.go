package bob

import (
    "strings"
    "unicode"
)

func areAllLettersCapitalized(s string) bool {
    hasLetters := false
    for _, r := range s {
        if unicode.IsLetter(r) {
            hasLetters = true
            if !unicode.IsUpper(r) {
                return false
            }
        }
    }
    return hasLetters
}

func isEffectivelyEmpty(s string) bool {
    return strings.TrimSpace(s) == ""
}

func isQuestion(s string) bool {
    return strings.HasSuffix(strings.TrimSpace(s), "?")
}

func Hey(remark string) string {
    remark = strings.TrimSpace(remark)

    switch {
    case isEffectivelyEmpty(remark):
        return "Fine. Be that way!"
    case areAllLettersCapitalized(remark) && isQuestion(remark):
        return "Calm down, I know what I'm doing!"
    case isQuestion(remark):
        return "Sure."
    case areAllLettersCapitalized(remark):
        return "Whoa, chill out!"
    default:
        return "Whatever."
    }
}