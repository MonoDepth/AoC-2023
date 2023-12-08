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

func mapTo[K comparable, V comparable](arr []K, callback func(item K) V) []V {
	var result []V

	for _, item := range arr {
		result = append(result, callback(item))
	}

	return result
}

func filter[K comparable](arr []K, callback func(item K) bool) []K {
	var result []K

	for _, item := range arr {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

func forEach[K comparable](arr []K, callback func(item K)) {
	for _, item := range arr {
		callback(item)
	}
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

func findOptimal(totalDurationMs, currentRecordMs int) (resultSet []int) {
	for i := 1; i < totalDurationMs; i++ {
		if calculateTravelDistance(totalDurationMs, i) > currentRecordMs {
			resultSet = append(resultSet, i)
		}
	}

	return resultSet
}

func calculateTravelDistance(totalDurationMs int, holdTimeMs int) int {
	return (totalDurationMs - holdTimeMs) * holdTimeMs
}
