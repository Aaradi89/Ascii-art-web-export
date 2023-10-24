package pkg

func IsNewLine(input string) bool {
	for i, x := range input {

		if i > 0 && x == 10 && input[i-1] == 13 { // There is a new line - 10 = Line Feed & 13 = Return
			return true
		}
	}
	return false
}
