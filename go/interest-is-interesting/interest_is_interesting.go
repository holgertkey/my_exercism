package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
    case balance < 0:
        return float32(3.213)
    case balance >= 0 && balance < 1000:
        return float32(0.5)
    case balance >= 1000 && balance < 5000:
        return float32(1.621)
    case balance >= 5000:
        return float32(2.475)
    default:
        return 0
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return float64(InterestRate(balance)) * balance / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    currentBalance := balance
    count := 0
    for currentBalance < targetBalance {
        currentBalance = AnnualBalanceUpdate(currentBalance)
        count++
    }
    return count
}
