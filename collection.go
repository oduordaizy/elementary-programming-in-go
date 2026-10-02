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


func LastWord(s string) string {
	if s == "" {
		return "\n"
	}

	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	if i < 0 {
		return "\n"
	}

	end := i

	for i >= 0 && s[i] != ' ' {
		i--
	}

	return s[i+1:end+1] + "\n"
}


func RepeatAlpha(s string) string {
	result := ""

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			for i := 0; i < int(r-'a')+1; i++ {
				result += string(r)
			}
		} else if r >= 'A' && r <= 'Z' {
			for i := 0; i < int(r-'A')+1; i++ {
				result += string(r)
			}
		} else {
			result += string(r)
		}
	}

	return result
}


func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

//cleanstr
func Cleanstr() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]
	var words []string
	var currentWord []byte

	// 2. Iterate through the string to extract words into a slice
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == ' ' || char == '\t' {
			// If we hit a space/tab and have collected a word, save it to the slice
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []byte{} // Reset for the next word
			}
		} else {
			// Append character to the current word
			currentWord = append(currentWord, char)
		}
	}

	// Catch the last word if the string doesn't end with whitespace
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	// 3. If there are no words to display, print a newline and exit
	if len(words) == 0 {
		fmt.Println()
		return
	}

	// 4. Build the final string by adding a single space between the words
	var result string
	for i := 0; i < len(words); i++ {
		result += words[i]
		if i < len(words)-1 {
			result += " "
		}
	}

	// 5. Print the result followed by a newline
	fmt.Println(result)
}


//expandstr
func Expandstr() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	word := false
	first := true

	for _, r := range s {
		if r == ' ' || r == '\t' {
			if word {
				word = false
			}
		} else {
			if !first && !word {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}

			z01.PrintRune(r)
			word = true
			first = false
		}
	}

	if !first {
		z01.PrintRune('\n')
	}
}

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	newWord := true

	for _, r := range s {
		if r == ' ' || r == '\t' {
			newWord = true
		} else {
			if newWord {
				if r >= 'a' && r <= 'z' {
					return false
				}
				newWord = false
			}
		}
	}

	return true
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	if n < 0 {
		result = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(rune('0'+digit)) + result
		n /= 10
	}

	return result
}

func Printrevcomb() {
	first := true

	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				if !first {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				z01.PrintRune(rune(k + '0'))

				first = false
			}
		}
	}

	z01.PrintRune('\n')
}

func ThirdTimeIsACharm(str string) string {
	result := ""

	for i, r := range str {
		if (i+1)%3 == 0 {
			result += string(r)
		}
	}

	return result + "\n"
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	count := 0

	for i, r := range str1 {
		found := false

		for j, r2 := range str1 {
			if i != j && r == r2 {
				found = true
				break
			}
		}

		if !found {
			for _, r2 := range str2 {
				if r == r2 {
					found = true
					break
				}
			}
		}

		if !found {
			count++
		}
	}

	for i, r := range str2 {
		found := false

		for j, r2 := range str2 {
			if i != j && r == r2 {
				found = true
				break
			}
		}

		if !found {
			for _, r2 := range str1 {
				if r == r2 {
					found = true
					break
				}
			}
		}

		if !found {
			count++
		}
	}

	return count
}

//Add Primesum
func main() {
	// Check if the number of arguments is not 1
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	// Convert the argument to an integer
	nb, err := strconv.Atoi(os.Args[1])
	// If it's not a valid integer or not a positive number/zero, display 0
	if err != nil || nb < 0 {
		fmt.Println(0)
		return
	}

	sum := 0
	// Sum all prime numbers less than or equal to nb
	for i := 2; i <= nb; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}

// Helper function to check if a number is prime
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

func CanJump(arr []uint) bool {
	// If the array is empty, return false
	if len(arr) == 0 {
		return false
	}

	curr := 0
	last := len(arr) - 1
	visited := make(map[int]bool)

	// Simulate the jumps step by step
	for curr < last {
		steps := int(arr[curr])

		// If we hit a 0 and haven't reached the end, we're stuck
		if steps == 0 {
			return false
		}

		// Prevent infinite loops if a cycle occurs
		if visited[curr] {
			return false
		}
		visited[curr] = true

		curr += steps

		// If the jump goes beyond the last index, it's invalid
		if curr > last {
			return false
		}
	}

	// Return true if we successfully landed on the last index
	return curr == last
}

func ConcatAlternate(slice1, slice2 []int) []int {
	var s1, s2 []int

	// Determine which slice is larger to start with it. 
	// If lengths are equal, slice1 takes precedence.
	if len(slice2) > len(slice1) {
		s1 = slice2
		s2 = slice1
	} else {
		s1 = slice1
		s2 = slice2
	}

	var result []int
	maxLength := len(s1)

	// Interleave elements from both slices
	for i := 0; i < maxLength; i++ {
		if i < len(s1) {
			result = append(result, s1[i])
		}
		if i < len(s2) {
			result = append(result, s2[i])
		}
	}

	return result
}


)

func fprime() {
	// Check if the number of arguments is not 1
	if len(os.Args) != 2 {
		return
	}

	// Parse the argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	// If invalid integer, less than or equal to 1 (since 1 has no prime factors), display nothing
	if err != nil || n <= 1 {
		return
	}

	var factors []int
	divisor := 2

	// Find all prime factors
	for n > 1 {
		if n%divisor == 0 {
			factors = append(factors, divisor)
			n /= divisor
		} else {
			divisor++
		}
	}

	// Display the factors separated by '*'
	for i := 0; i < len(factors); i++ {
		fmt.Printf("%d", factors[i])
		if i < len(factors)-1 {
			fmt.Print("*")
		}
	}
	fmt.Println()
}




func Hiddenp() {
	// Check if the number of arguments is not 2
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// If s1 is an empty string, it is considered hidden in any string
	if s1 == "" {
		fmt.Println(1)
		return
	}

	i := 0 // Index for tracking characters in s1
	for j := 0; j < len(s2); j++ {
		// If current character in s2 matches the expected character in s1
		if s2[j] == s1[i] {
			i++
			// If we have matched all characters of s1, it is hidden
			if i == len(s1) {
				fmt.Println(1)
				return
			}
		}
	}

	// If we finished s2 without matching all characters of s1
	fmt.Println(0)
}


func Reversesstrcap() {
	// If there are no arguments (only os.Args[0] exists), display nothing
	if len(os.Args) < 2 {
		return
	}

	// Process each argument provided to the program
	for argIdx := 1; argIdx < len(os.Args); argIdx++ {
		s := os.Args[argIdx]
		runes := []rune(s)
		
		for i := 0; i < len(runes); i++ {
			// A word ends if it's the last character of the string 
			// or if the next character is a space or tab.
			isLastOfWord := false
			if i == len(runes)-1 {
				isLastOfWord = true
			} else if runes[i+1] == ' ' || runes[i+1] == '\t' {
				isLastOfWord = true
			}

			// Apply case conversion
			if isLastOfWord {
				// Convert to uppercase if it's a lowercase letter
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] -= 32
				}
			} else {
				// Convert to lowercase if it's an uppercase letter
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] += 32
				}
			}
		}

		fmt.Println(string(runes))
	}
}


func wdmatch() {
	if len(os.Args) != 3 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]

	j := 0

	for i := 0; i < len(second) && j < len(first); i++ {
		if second[i] == first[j] {
			j++
		}
	}

	if j == len(first) {
		for _, char := range first {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	dot := -1
	digits := 0

	for i, r := range dec {
		if r == '.' {
			if dot != -1 {
				return dec + "\n"
			}
			dot = i
		} else if r >= '0' && r <= '9' {
			digits++
		} else if r == '-' && i == 0 {
			continue
		} else {
			return dec + "\n"
		}
	}

	if dot == -1 {
		return dec + "\n"
	}

	afterDot := dec[dot+1:]

	// If everything after the decimal point is 0,
	// return the original number.
	allZero := true
	for _, r := range afterDot {
		if r != '0' {
			allZero = false
			break
		}
	}

	if allZero {
		return dec + "\n"
	}

	// Remove the decimal point.
	result := ""
	for _, r := range dec {
		if r != '.' {
			result += string(r)
		}
	}

	return result + "\n"
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	n1 := len(slice1)
	n2 := len(slice2)

	if n1 >= n2 {
		// slice1 is larger or equal
		// First, add the excess elements from the back of slice1 in reverse order
		for i := n1 - 1; i >= n2; i-- {
			result = append(result, slice1[i])
		}
		// Then, interleave the remaining equal-sized portions starting with slice1
		for i := n2 - 1; i >= 0; i-- {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	} else {
		// slice2 is larger
		// First, add the excess elements from the back of slice2 in reverse order
		for i := n2 - 1; i >= n1; i-- {
			result = append(result, slice2[i])
		}
		// Then, interleave the remaining equal-sized portions starting with slice1
		for i := n1 - 1; i >= 0; i-- {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	}

	return result
}


func Slice(a []string, nbrs ...int) []string {
	length := len(a)
	if length == 0 || len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := length

	// Handle negative start index
	if start < 0 {
		start = length + start
		if start < 0 {
			start = 0
		}
	}

	// If a second integer is provided
	if len(nbrs) > 1 {
		end = nbrs[1]
		// Handle negative end index
		if end < 0 {
			end = length + end
			if end < 0 {
				end = 0
			}
		}
	}

	// Boundary checks and invalid range cases (e.g., start > end)
	if start > length {
		return nil
	}
	if end > length {
		end = length
	}
	if start > end {
		return nil
	}

	return a[start:end]
}


func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Remove spaces
	var cleaned []rune

	for _, char := range str {
		if char != ' ' {
			cleaned = append(cleaned, char)
		}
	}

	// Less than 5 characters
	if len(cleaned) < 5 {
		return "Invalid Input\n"
	}

	var result []rune
	count := 0

	for i := 0; i < len(cleaned); i++ {
		result = append(result, cleaned[i])
		count++

		if count == 5 {
			count = 0

			// If there is another character, add a space
			// and skip the next character.
			if i+1 < len(cleaned) {
				result = append(result, ' ')
				i++
			}
		}
	}

	return string(result) + "\n"
}

