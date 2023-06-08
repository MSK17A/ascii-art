package asciiART

func ContainsLetter(letters string, input_letter string) bool {
	for _, letter := range letters {
		if letter == rune(input_letter[0]) {
			return true
		}
	}

	return false
}
