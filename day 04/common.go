package main

func selectNotEmpty(arr []string, callback func(string) string) []string {
	var result []string

	for _, item := range arr {
		itm := callback(item)
		if itm != "" {
			result = append(result, itm)
		}
	}

	return result
}

func pow(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}
