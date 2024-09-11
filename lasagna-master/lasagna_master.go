package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timeForLayer int) int {
	if timeForLayer == 0 {
		timeForLayer = 2
	}
	return len(layers) * timeForLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (grams int, liters float64) {
	for _, v := range layers {
		if v == "noodles" {
			grams += 50
		} else if v == "sauce" {
			liters += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaled []float64
	for i := range quantities {
		scaled = append(scaled, quantities[i]/2*float64(portions))
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
