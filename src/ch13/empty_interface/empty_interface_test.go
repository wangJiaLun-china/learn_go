package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i,ok:=p.(int); ok{
	//	fmt.Println("Integer", i)
	//}
	//if s,ok:=p.(string); ok{
	//	fmt.Println("string", s)
	//	return
	//}
	//fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

// Go 接口 倾向于使用小的接口定义，很多接口只包含一个方法
// 较大的接口，可以通过较小的接口组合而成
// 只依赖于必要功能的最小接口
func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
