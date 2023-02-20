package hub

var valueSet = []float64{-1, 0, 0.5, 1, 2, 3, 5, 8, 13, 21}

func IsValidValue(value float64) bool {
	for _, v := range valueSet {
		if v == value {
			return true
		}
	}

	return false
}
