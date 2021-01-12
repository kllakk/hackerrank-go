package tests

import (
	"bufio"
	"fmt"
	fs "hackerrank-go/src/find-strings"
	"hackerrank-go/tools"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func execFindStrings(testId int, ts *testing.T) {
	go tools.PanicOnTimeout(4 * time.Second)

	file, err := os.Open(fmt.Sprintf("find-strings/test-%d-input.txt", testId))
	tools.CheckError(err)
	defer file.Close()

	reader := bufio.NewReaderSize(file, tools.MaxBuffer)

	wCount, err := strconv.ParseInt(tools.ReadLine(reader), 10, 64)
	tools.CheckError(err)

	var w []string

	for wItr := 0; wItr < int(wCount); wItr++ {
		wItem := tools.ReadLine(reader)
		w = append(w, wItem)
	}

	queriesCount, err := strconv.ParseInt(tools.ReadLine(reader), 10, 64)
	tools.CheckError(err)

	var queries []int32

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(tools.ReadLine(reader), 10, 64)
		tools.CheckError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := fs.FindStrings(w, queries)

	output := tools.ReadAsSlices(fmt.Sprintf("find-strings/test-%d-output.txt", testId))
	if !reflect.DeepEqual(result, output) {
		ts.Errorf("Expected %v, got %v", output, result)
	}
}

func TestFindStrings00(ts *testing.T) {
	execFindStrings(0, ts)
}

func TestFindStrings01(ts *testing.T) {
	execFindStrings(1, ts)
}

func TestFindStrings02(ts *testing.T) {
	execFindStrings(2, ts)
}

func TestFindStrings03(ts *testing.T) {
	execFindStrings(3, ts)
}

func TestFindStrings04(ts *testing.T) {
	execFindStrings(4, ts)
}

func TestFindStrings05(ts *testing.T) {
	execFindStrings(5, ts)
}

func TestFindStrings06(ts *testing.T) {
	execFindStrings(6, ts)
}

func TestFindStrings07(ts *testing.T) {
	execFindStrings(7, ts)
}
