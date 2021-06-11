package main

func SumAllTails(numbersToSum ...[]int) []int {
	// Initialize a slice
	var sums []int

	// Loop over the supplied slice using "range"
	for _, numbers := range numbersToSum {
		// Slice the slice to grab the tail.
		// This should grab all but the first
		// member of the slice.
		tail := numbers[1:]
		// Append values to "sums". Appended values are
		// what was in sums (nothing) and the sum of 
		// the slice tail.
		sums = append(sums, Sum(tail))
	}

	return sums
}