package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	// create var to store sum
	var sum int

	// iterate each value of numbers
	// add value to var
	for _, val := range numbers {
		sum += val
	}
	return sum
}
