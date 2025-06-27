package cipher

import (
    "strings"
)

type shift struct {
    distance int
}

type vigenere struct {
    key string
}

func NewCaesar() Cipher {
    return NewShift(3)
}

func NewShift(distance int) Cipher {
    if distance <= -26 || distance >= 26 || distance == 0 {
        return nil
    }
    return shift{distance: distance}
}

func NewVigenere(key string) Cipher {
    if key == "" {
        return nil
    }

    hasNonA := false
    for _, r := range key {
        if r < 'a' || r > 'z' {
            return nil
        }
        if r != 'a' {
            hasNonA = true
        }
    }

    if !hasNonA {
        return nil
    }

    return vigenere{key: key}
}

func (c shift) Encode(input string) string {
    return shiftString(input, c.distance)
}

func (c shift) Decode(input string) string {
    return shiftString(input, -c.distance)
}

func shiftString(s string, distance int) string {
    var result strings.Builder

    for _, r := range s {
        if r >= 'A' && r <= 'Z' {
            r = r + 'a' - 'A'
        }
        if r >= 'a' && r <= 'z' {
            r = 'a' + (r - 'a' + rune(distance) + 26) % 26
            result.WriteRune(r)
        }
    }

    return result.String()
}

func (v vigenere) Encode(input string) string {
    return v.cipher(input, true)
}

func (v vigenere) Decode(input string) string {
    return v.cipher(input, false)
}

func (v vigenere) cipher(input string, encode bool) string {
    var result strings.Builder
    keyRunes := []rune(v.key)
    keyIndex := 0

    for _, r := range input {
        if r >= 'A' && r <= 'Z' {
            r = r + 'a' - 'A'
        }
        if r >= 'a' && r <= 'z' {
            shift := int(keyRunes[keyIndex] - 'a')
            if !encode {
                shift = -shift
            }
            r = 'a' + (r - 'a' + rune(shift) + 26) % 26
            result.WriteRune(r)
            keyIndex = (keyIndex + 1) % len(keyRunes)
        }
    }

    return result.String()
}
