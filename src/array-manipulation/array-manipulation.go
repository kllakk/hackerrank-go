package array_manipulation

func ArrayManipulation(n int32, queries [][]int32) int64 {
	var max int64 = 0

	var previous = make([]int64, n)
	for i := range previous {
		previous[i] = 0
	}

	for q := range queries {
		query := queries[q]
		chunk := previous[query[0] - 1:query[1]]

		for i := range chunk {
			chunk[i] = chunk[i] + int64(query[2])

			if chunk[i] > max {
				max = chunk[i]
			}
		}
	}

	return max
}