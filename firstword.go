package piscine

func FisrtWord(s string) string {

	if s == "" {
		return "\n"
	}

	for i = 0; i < len(s); i++ {
	if s[i] == " "{
		return s[:i] + "\n"
	}

}