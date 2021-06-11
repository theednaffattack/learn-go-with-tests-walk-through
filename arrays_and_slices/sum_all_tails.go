package main

func SumAllTails(numbersToSum ...[]int) []int {
	// Create a slice
	var sums []int

	// Loop over the supplied slice using "range"
	for _, numbers := range numbersToSum {
		// If there is no tail, append zero
		if len(numbers) == 0 {
			sums = append(sums, 0)
			// Else create the tail and append its
			// sum.
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}