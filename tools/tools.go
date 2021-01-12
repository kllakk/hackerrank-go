package tools

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"
)

const MaxBuffer = 1024 * 1024

func PanicOnTimeout(d time.Duration) {
	<-time.After(d)
	panic("Test timed out")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadAsSlices(filepath string) []string {
	buffer := make([]byte, MaxBuffer)

	var result []string
	file, err := os.Open(filepath)
	CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buffer, MaxBuffer)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	err = scanner.Err()
	CheckError(err)

	return result
}

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}