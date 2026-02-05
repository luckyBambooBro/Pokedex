package main

func cleanInput(text string) []string {
	var words []string
	var word string
	for _, char := range text {
		if string(char) != " " {
			word += string(char)
		}
		if string(char) == " " && word != " " {
			words = append(words, word)
		}
	}
	return words
}