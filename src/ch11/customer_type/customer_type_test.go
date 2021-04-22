package customer_type

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 自定义类型
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 函数式编程
// 推荐阅读 计算机程序的构造和解释
func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}
