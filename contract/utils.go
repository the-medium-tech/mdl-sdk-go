package contract

func StringArrayToTwoDimensionalArray(args []string) [][]byte {
	bytes := make([][]byte, len(args))
	for i, v := range args {
		bytes[i] = []byte(v)
	}
	return bytes
}
