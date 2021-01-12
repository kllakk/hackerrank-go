package tests

import (
	"bufio"
	"fmt"
	"hackerrank-go/tools"
	"os"
	quts "queue-using-two-stacks"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func execQueueUsingTwoStacks(testId int, ts *testing.T) {
	go tools.PanicOnTimeout(4 * time.Second)

	file, err := os.Open(fmt.Sprintf("queue-using-two-stacks/test-%d-input.txt", testId))
	tools.CheckError(err)
	defer file.Close()

	reader := bufio.NewReaderSize(file, tools.MaxBuffer)

	qCount, err := strconv.ParseInt(tools.ReadLine(reader), 10, 64)
	tools.CheckError(err)

	var queries [][]string
	for qItr := 0; qItr < int(qCount); qItr++ {
		qItem := tools.ReadLine(reader)
		queries = append(queries, strings.Split(qItem, " "))
	}

	result := quts.QueueUsingTwoStacks(queries)

	output := tools.ReadAsSlices(fmt.Sprintf("queue-using-two-stacks/test-%d-output.txt", testId))
	if !reflect.DeepEqual(result, output) {
		ts.Errorf("Expected %v, got %v", output, result)
	}
}

func TestQueueUsingTwoStacks00(ts *testing.T) {
	execQueueUsingTwoStacks(0, ts)
}

func TestQueueUsingTwoStacks01(ts *testing.T) {
	execQueueUsingTwoStacks(1, ts)
}

func TestQueueUsingTwoStacks15(ts *testing.T) {
	execQueueUsingTwoStacks(15, ts)
}