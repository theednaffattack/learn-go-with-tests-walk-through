package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for index := 0; index < repeatCount; index++ {
		repeated += character
	}
	return repeated
}