package lsproduct

import (
    "errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if span < 0 {
        return 0, errors.New("span must not be negative")
    }
    if span > len(digits) {
        return 0, errors.New("span must not be bigger than string length")
    }

    var maxProduct int64 = 0

    for i := 0; i <= len(digits) - span; i++ {
        product := int64(1)
        for j := 0; j < span; j++ {
            d := digits[i + j]
            if d < '0' || d > '9' {
                return 0, errors.New("input must only contain digits")
            }
            product *= int64(d - '0')
        }
        if product > maxProduct {
            maxProduct = product
        }
    }

    return maxProduct, nil
}
