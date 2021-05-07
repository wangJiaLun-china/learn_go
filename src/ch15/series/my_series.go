package series

import "fmt"

func init()  {
	fmt.Println("init2")
}

func init()  {
	fmt.Println("init1")
}

func GetNum() int {
	return 10
}

func getNum() int {
	return 11
}
