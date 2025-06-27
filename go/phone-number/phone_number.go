package phonenumber

import (
    "errors"
    "regexp"
)

func Number(phoneNumber string) (string, error) {
    re := regexp.MustCompile(`[0-9]`)
    digits := re.FindAllString(phoneNumber, -1)

    if len(digits) == 11 && digits[0] == "1" {
        digits = digits[1:]
    }

    if len(digits) != 10 {
        return "", errors.New("invalid phone number")
    }

    if digits[0] < "2" || digits[3] < "2" {
        return "", errors.New("invalid area or exchange code")
    }

    return join(digits), nil
}

func AreaCode(phoneNumber string) (string, error) {
    num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
    num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return "(" + num[:3] + ") " + num[3:6] + "-" + num[6:], nil
}

func join(digits []string) string {
    var result string
    for _, d := range digits {
        result += d
    }
    return result
}
