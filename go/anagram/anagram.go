package anagram

import (
    "strings"
    "unicode"
    "sort"
)

func normalize(s string) string {
    var letters []rune
    for _, r := range s {
        if unicode.IsLetter(r) {
            letters = append(letters, unicode.ToLower(r))
        }
    }
    sort.Slice(letters, func(i, j int) bool {
        return letters[i] < letters[j]
    })
    return string(letters)
}

func isAnagram(subject, candidate string) bool {
    if strings.EqualFold(subject, candidate) {
        return false
    }
    return normalize(subject) == normalize(candidate)
}

func Detect(subject string, candidates []string) []string {
    var anagrams []string
    for _, candidate := range candidates {
        if isAnagram(subject, candidate) {
            anagrams = append(anagrams, candidate)
        }
    }
    return anagrams
}