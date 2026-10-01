package elementaryprogrammingingo

func HashCode(dec string) string {
	length := len(dec)
	var result []byte

	for i := 0; i < length; i++ {
		hash := (int(dec[i]) + length) % 127

		if hash < 32 {
			hash += 33
		}

		result = append(result, byte(hash))
	}
	return string(result)
}
