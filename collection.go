func PrintMemory(arr [10]byte) {
	// 1. Print the hexadecimal representation in chunks of 4 bytes per line
	for i := 0; i < len(arr); i += 4 {
		end := i + 4
		if end > len(arr) {
			end = len(arr)
		}

		for j := i; j < end; j++ {
			fmt.Printf("%02x", arr[j])
			if j < end-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// 2. Print the ASCII graphic characters (non-printable characters become '.')
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func ConcatSlice(slice1, slice2 []int) []int {
	return append(slice1, slice2...)
}

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result []byte

	count := 1

	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
		} else {
			result = append(result, byte('0'+count))
			result = append(result, s[i-1])
			count = 1
		}
	}

	return string(result)
}

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	// Loop backwards from nb down to 2
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}

	return 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}


func FromTo(from int, to int) string {
	// Check if any argument is out of bounds (< 0 or > 99)
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string

	// Handle ascending order (from <= to)
	if from <= to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0"
			}
			result += fmt.Sprintf("%d", i)
			if i != to {
				result += ", "
			}
		}
	} else {
		// Handle descending order (from > to)
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0"
			}
			result += fmt.Sprintf("%d", i)
			if i != to {
				result += ", "
			}
		}
	}

	result += "\n"
	return result
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Validate camelCase
	for i := 0; i < len(s); i++ {
		c := s[i]

		// Only letters are allowed
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return s
		}

		// Cannot end with a capital letter
		if i == len(s)-1 && c >= 'A' && c <= 'Z' {
			return s
		}

		// Two capital letters cannot be next to each other
		if i > 0 &&
			c >= 'A' && c <= 'Z' &&
			s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}
	}

	// Convert to snake_case
	result := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				result += "_"
			}

			c += 'a' - 'A'
		}

		result += string(c)
	}

	return result
}

