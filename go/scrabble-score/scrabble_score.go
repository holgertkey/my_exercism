package scrabble

import "strings"

func Score(word string) int {
    lWord := strings.ToLower(word)
    lettersValues := map[string]int{
        "aeioulnrst": 1,
        "dg":         2,
        "bcmp":       3,
        "fhvwy":      4,
        "k":          5,
        "jx":         8,
        "qz":        10,
    }
    count := 0
    for _, letter := range lWord {
        for key, value := range lettersValues {
            if strings.Contains(key, string(letter)) {
                count += value
                break
            }
        }
    }
    return count    
}
