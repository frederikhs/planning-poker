package hub

var valueSet = []int{-1, 0, 1, 2, 3, 5, 8, 13, 21}

func IsValidValue(value int) bool {
	for _, v := range valueSet {
		if v == value {
			return true
		}
	}

	return false
}
