package iteration

// const repeatCount = 5

func Repeat(character string, iterations int) string {
	var repeated string
	for index := 0; index < iterations; index++ {
		repeated += character
	}
	return repeated
}