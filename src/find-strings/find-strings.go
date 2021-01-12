package find_strings

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

func sliceContainsSubstring(slice []string, str string) bool {
	result := false

	for _, s := range slice {
		if strings.Contains(s, str) {
			result = true
			break
		}
	}

	return result
}

func md5String(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

func sliceUnique(slice []string) []string {
	keys := make(map[string]bool)

	var result []string

	for _, item := range slice {
		hash := md5String(item)
		if _, value := keys[hash]; !value {
			keys[hash] = true
			result = append(result, item)
		}
	}

	return result
}

func queriesIndexes(sequence []string, queries []int32) []string {
	var result []string

	for _, query := range queries {
		if int(query - 1) >= len(sequence) || int(query - 1) < 0 {
			result = append(result, "INVALID")
		} else {
			result = append(result, sequence[query - 1])
		}
	}

	return result
}

func substringSequence(substring string, keys map[string]bool) []string {

	var result []string
	for i := 0; i < len(substring); i++ {

		size := len(substring) - i

		for j := 0; j <= len(substring) - size; j++ {
			sequence := substring[j:size + j]

			hash := md5String(sequence)
			if _, value := keys[hash]; !value {
				keys[hash] = true
				result = append(result, sequence)
			}
		}
	}

	return result
}

/*
 * Complete the findStrings function below.
 */
func FindStrings(w []string, queries []int32) []string {
	/*
	 * Write your code here.
	 */

	var wUnique []string

	// длинные строки выше
	sort.Slice(w, func (i, j int) bool {
		return len(w[i]) > len(w[j])
	})

	for _, str := range w {
		if !sliceContainsSubstring(wUnique, str) {
			wUnique = append(wUnique, str)
		}
	}

	var sequences []string
	keys := make(map[string]bool)

	for _, sequence := range wUnique {
		// abc substrings[0] a, b, c, ab, bc, abc
		// cd substringsS[1] c, d, cd
		sequences = append(sequences, substringSequence(sequence, keys)...)
	}

	sort.Strings(sequences)

	// 1 a, 5 bc, 7 cd
	result := queriesIndexes(sequences, queries)

	return result
}