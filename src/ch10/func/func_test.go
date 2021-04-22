package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
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

// 可变参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// defer
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
