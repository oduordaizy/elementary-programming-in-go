package elementaryprogrammingingo

import "fmt"

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
