package solve_me_first

import (
	"fmt"
	"os"
	"strconv"
)

func SolveMeFirst(a uint32,b uint32) uint32{
	return a + b
}

func main() {
	var a, b, res uint32

	if a, err := strconv.ParseUint(os.Args[1], 10, 32); err == nil {
		fmt.Printf("%T, %v\n", a, a)
	}

	if b, err := strconv.ParseUint(os.Args[2], 10, 32); err == nil {
		fmt.Printf("%T, %v\n", b, b)
	}

	res = SolveMeFirst(a,b)
	fmt.Println(res)
}