package iteration

func Repeat(character string) string {
	var repeated string
	for index := 0; index < 5; index++ {
		repeated = repeated + character
	}
	return repeated
}