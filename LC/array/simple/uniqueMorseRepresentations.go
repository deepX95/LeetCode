package simple

// uniqueMorseRepresentations
func uniqueMorseRepresentations(words []string) int {
	morses := []string{
		".-", "-...", "-.-.", "-..", ".",
		"..-.", "--.", "....", "..", ".---",
		"-.-", ".-..", "--", "-.", "---",
		".--.", "--.-", ".-.", "...", "-",
		"..-", "...-", ".--", "-..-", "-.--",
		"--..",
	}
	dst := map[string]struct{}{}

	for _, word := range words {
		morse := ""
		for _, ch := range word {
			morse += morses[ch-'a']
		}
		dst[morse] = struct{}{}
	}
	return len(dst)
}
