package tests

import (
	"bufio"
	"fmt"
	saveHumanity "hackerrank-go/src/save-humanity"
	"hackerrank-go/tools"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func execSaveHumanity(testId int, ts *testing.T) {
	go tools.PanicOnTimeout(4 * time.Second)

	file, err := os.Open(fmt.Sprintf("save-humanity/test-%d-input.txt", testId))
	tools.CheckError(err)
	defer file.Close()

	var result []string

	var t int32 = 0
	fistLine := true

	buffer := make([]byte, tools.MaxBuffer)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buffer, tools.MaxBuffer)

	for scanner.Scan() {
		if fistLine {
			tTemp, err := strconv.ParseInt(scanner.Text(), 10, 64)
			tools.CheckError(err)
			t = int32(tTemp)
			fistLine = false
		} else {
			if t > 0 {

				pv := strings.Split(scanner.Text(), " ")
				p := pv[0]
				v := pv[1]

				result = append(result, saveHumanity.VirusIndices(p, v))

				t--
			} else {
				break
			}
		}
	}
	err = scanner.Err()
	tools.CheckError(err)

	output := tools.ReadAsSlices(fmt.Sprintf("save-humanity/test-%d-output.txt", testId))
	if !reflect.DeepEqual(result, output) {
		ts.Errorf("Expected %v, got %v", output, result)
	}
}

func TestSaveHumanity00(ts *testing.T) {
	execSaveHumanity(0, ts)
}

func TestSaveHumanity01(ts *testing.T) {
	execSaveHumanity(1, ts)
}

func TestSaveHumanity03(ts *testing.T) {
	execSaveHumanity(3, ts)
}

func TestSaveHumanity05(ts *testing.T) {
	execSaveHumanity(5, ts)
}