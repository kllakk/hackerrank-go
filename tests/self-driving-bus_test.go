package tests

import (
	"bufio"
	"fmt"
	module "hackerrank-go/src/self-driving-bus"
	"hackerrank-go/tools"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func execSelfDrivingBus00(testId int, ts *testing.T) {
	//go tools.PanicOnTimeout(4 * time.Second)

	file, err := os.Open(fmt.Sprintf("self-driving-bus/test-%d-input.txt", testId))
	tools.CheckError(err)
	defer file.Close()

	buffer := make([]byte, tools.MaxBuffer)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buffer, tools.MaxBuffer)

	fistLine := true
	var tree [][]int32
	var n int32
	for scanner.Scan() {
		if fistLine {
			nTemp, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
			tools.CheckError(err)
			n = int32(nTemp)
			fistLine = false
		} else {
			for i := 0; i < int(n) - 1; i++ {
				treeRowTemp := strings.Split(strings.TrimRight(scanner.Text()," \t\r\n"), " ")

				var treeRow []int32
				for _, treeRowItem := range treeRowTemp {
					treeItemTemp, err := strconv.ParseInt(treeRowItem, 10, 64)
					tools.CheckError(err)
					treeItem := int32(treeItemTemp)
					treeRow = append(treeRow, treeItem)
				}

				if len(treeRow) != 2 {
					panic("Bad input")
				}

				tree = append(tree, treeRow)
			}
		}
	}
	err = scanner.Err()
	tools.CheckError(err)

	result := module.Solve(tree)

	output := tools.ReadAsSlices(fmt.Sprintf("self-driving-bus/test-%d-output.txt", testId))
	if !reflect.DeepEqual(result, output) {
		ts.Errorf("Expected %v, got %v", output, result)
	}
}

func TestSelfDrivingBus00(ts *testing.T) {
	execSelfDrivingBus00(0, ts)
}

func TestSelfDrivingBus01(ts *testing.T) {
	execSelfDrivingBus00(1, ts)
}

func TestSelfDrivingBus06(ts *testing.T) {
	execSelfDrivingBus00(6, ts)
}

//func TestSelfDrivingBus34(ts *testing.T) {
//	execSelfDrivingBus00(34, ts)
//}
