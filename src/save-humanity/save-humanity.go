package save_humanity

import (
	"fmt"
	"strconv"
	"strings"
)

func compare(p string, v string) bool {
	if p == v || len(v) == 1 {
		return true
	}

	// одна из половин должна совпадать!
	chunk := len(p) / 2

	pFirst := p[:chunk]
	pLast := p[chunk:]

	vFirst := v[:chunk]
	vLast := v[chunk:]

	if pFirst == vFirst {
		return compare(pLast, vLast)
	} else if pLast == vLast {
		return compare(pFirst, vFirst)
	} else {
		return false
	}
}

func VirusIndices(p string, v string) string {
	/*
	 * Print the answer for this test case in a single line
	 */
	var indexes []string

	for i:= 0; i <= len(p) - len(v); i++ {
		if compare(p[i:i + len(v)], v) {
			indexes = append(indexes, strconv.Itoa(i))
		}
	}

	if len(indexes) == 0 {
		indexes = append(indexes, "No Match!")
	}

	fmt.Println(strings.Join(indexes, " "))
	return strings.Join(indexes, " ")
}
