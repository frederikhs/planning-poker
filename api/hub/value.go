package hub

var valueSet = []int{0, 1, 3, 5, 8, 13, 21}

func IsValidValue(value int) bool {
	for _, v := range valueSet {
		if v == value {
			return true
		}
	}

	return false
}
