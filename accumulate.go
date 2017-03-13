package accumulate

const testVersion = 1

func Accumulate(given []string, converter func(string) string) []string {
	n := len(given)
	result := make([]string, n)
	for i, g := range given {
		result[i] = converter(g)
	}
	return result
}
