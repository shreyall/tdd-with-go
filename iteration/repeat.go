package iteration

// Repeat returns a string of 5 characters
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}

	return repeated
}
