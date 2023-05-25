package inkwell

func IsVowelBegin(input string) bool {
	ar := []rune(input)

	if ar[0] == 'a' || ar[0] == 'A' || ar[0] == 'e' || ar[0] == 'E' || ar[0] == 'i' || ar[0] == 'I' || ar[0] == 'o' || ar[0] == 'O' || ar[0] == 'u' || ar[0] == 'U' || ar[0] == 'h' || ar[0] == 'H' {
		return true
	}

	return false
}
