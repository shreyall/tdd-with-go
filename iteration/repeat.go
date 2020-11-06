package iteration

// Repeat returns a string of 5 characters
func Repeat(character string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated = repeated + character
	}

	return repeated
}
