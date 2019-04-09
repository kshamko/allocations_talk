package example4

func sliceAppend(count int) []int {
	out := []int{}
	for j := 0; j < count; j++ {
		out = append(out, j)
	}

	return out
}

func sliceIndex(count int) []int {
	out := make([]int, count)
	for j := 0; j < count; j++ {
		out[j] = j
	}
	return out
}
