package thefarm

import (
    "fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    famount, err := fc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }
    factor, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    return famount * factor / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    cows int
    message string
}

func (i InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", i.cows, i.message)
}

func ValidateNumberOfCows(cows int) error {
    switch {
    case cows < 0:
        return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
    case cows == 0:
        return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
    default:
        return nil
    }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
