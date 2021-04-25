package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	fmt.Println("测试执行")

	//os.Exit(-1)
}

func TestDeferDemo(t *testing.T) {
	defer fmt.Println("111")
	g(0)
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
