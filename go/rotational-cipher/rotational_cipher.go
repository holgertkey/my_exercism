package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
    result := []rune{}

    for _, r := range plain {
        switch {
        case r >= 'A' && r <= 'Z':
            shifted := ((r - 'A') + rune(shiftKey)) % 26
            result = append(result, 'A' + shifted)
        case r >= 'a' && r <= 'z':
            shifted := ((r - 'a') + rune(shiftKey)) % 26
            result = append(result, 'a' + shifted)
        default:
            result = append(result, r)
        }
    }

    return string(result)
}
