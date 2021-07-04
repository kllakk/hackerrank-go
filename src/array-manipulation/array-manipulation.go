package array_manipulation

func ArrayManipulation(n int32, queries [][]int32) int64 {
	values := make([]int64, n)
	for i := range queries {
		a := int64(queries[i][0])
		b := int64(queries[i][1])
		k := int64(queries[i][2])
		values[a-1] += k
		if b < int64(n) {
			values[b] -= k
		}
	}

	var max int64 = 0
	var sum int64 = 0
	for i := range values {
		sum = sum + values[i]
		if max < sum {
			max = sum
		}
	}

	return max
}