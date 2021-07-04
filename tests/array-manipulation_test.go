package tests

import (
	"bufio"
	"fmt"
	module "hackerrank-go/src/array-manipulation"
	"hackerrank-go/tools"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func execArrayManipulation00(testId int, ts *testing.T) {
	//go tools.PanicOnTimeout(4 * time.Second)

	file, err := os.Open(fmt.Sprintf("array-manipulation/test-%d-input.txt", testId))
	tools.CheckError(err)
	defer file.Close()

	var result []string

	buffer := make([]byte, tools.MaxBuffer)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buffer, tools.MaxBuffer)

	fistLine := true
	var queries [][]int32
	var n int32
	var m int
	for scanner.Scan() {
		if fistLine {
			firstMultipleInput := strings.Split(strings.TrimSpace(scanner.Text()), " ")
			tools.CheckError(err)

			nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
			tools.CheckError(err)
			n = int32(nTemp)

			mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
			tools.CheckError(err)
			m = int(mTemp)

			fistLine = false
		} else {
			queriesRowTemp := strings.Split(strings.TrimRight(scanner.Text()," \t\r\n"), " ")
			tools.CheckError(err)

			var queriesRow []int32
			for _, queriesRowItem := range queriesRowTemp {
				queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
				tools.CheckError(err)
				queriesItem := int32(queriesItemTemp)
				queriesRow = append(queriesRow, queriesItem)
			}

			if len(queriesRow) != 3 {
				panic("Bad input")
			}

			queries = append(queries, queriesRow)
		}
	}
	err = scanner.Err()
	tools.CheckError(err)

	if m != len(queries) {
		panic("Wrong queries")
	}

	resultInt := module.ArrayManipulation(n, queries)
	result = append(result, strconv.FormatInt(resultInt, 10))

	output := tools.ReadAsSlices(fmt.Sprintf("array-manipulation/test-%d-output.txt", testId))
	if !reflect.DeepEqual(result, output) {
		ts.Errorf("Expected %v, got %v", output, result)
	}
}

func TestArrayManipulation00(ts *testing.T) {
	execArrayManipulation00(0, ts)
}

func TestArrayManipulation01(ts *testing.T) {
	execArrayManipulation00(1, ts)
}

func TestArrayManipulation02(ts *testing.T) {
	execArrayManipulation00(2, ts)
}