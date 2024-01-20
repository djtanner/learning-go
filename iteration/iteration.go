package iteration

func Repeat(letter string, repeatCount int) string {
	//declares and initializes
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += letter
	}
	return repeated
}
