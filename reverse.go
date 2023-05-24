package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	// create var represeting new string
	var reverse string

	// iterate backwards over input
	for i := len(word) - 1; i >= 0; i-- {
		// add each rune to new string
		rune := rune(word[i])
		reverse += string(rune)
	}
	return reverse
}
