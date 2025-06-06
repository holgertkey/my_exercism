package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        return len(layers) * 2
    } else {
        return len(layers) * time
    }    
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (noodles int, sauce float64) {
    noodles = 0
    sauce = 0.0
    for _, value := range ingredients {
        if value == "noodles" {
            noodles += 50
        }
        if value == "sauce" {
            sauce += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    if len(friendsList) > 0 && len(myList) > 0 {
        myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
    }
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    newQuantities := make([]float64, len(quantities))
    for i, val := range quantities {
        newQuantities[i] = val / 2 * float64(portions)
    }
    return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
