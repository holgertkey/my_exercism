package pangram

import "strings"

func IsPangram(input string) bool {
    str := strings.ToLower(input)

    for ch := 'a'; ch <= 'z'; ch++ {
        if !strings.ContainsRune(str, ch) {
            return false
        }        
    }
    
    return true
}
