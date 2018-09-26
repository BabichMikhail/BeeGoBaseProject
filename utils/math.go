package utils

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CompareUint8(v1, v2 uint8) int8 {
	if v1 > v2 {
		return 1
	}
	if v2 > v1 {
		return -1
	}
	return 0
}

func AbsInt(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func PowInt(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result = result * a
	}
	return result
}

func Sign(val int) int {
	if val == 0 {
		return 0
	}
	return val / AbsInt(val)
}
