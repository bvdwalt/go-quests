package loops

import "strings"

func SumEvenNumbers(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

func KeepOnlyConsonants(strs []string) []string {
	res := []string{}
	for _, s := range strs {
		strResult := ""
		for _, char := range s {
			switch strings.ToLower(string(char)) {
			case "a", "e", "i", "o", "u":
				continue
			default:
				strResult += string(char)
			}
		}
		if strResult != "" {
			res = append(res, strResult)
		}
	}
	return res
}
