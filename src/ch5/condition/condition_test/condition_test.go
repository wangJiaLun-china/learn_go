package condition_test

import "testing"

func TestIfMultiSec(t *testing.T)  {

	if a:= 1 ==1;a {
		t.Log("1")
	}
	// if 支持两段式写法；可以变量赋值
	//if v,err := someFun(); err==nil{
	//	t.Log(v)
	//}else {
	//	t.Log(err)
	//}
}


func TestSwitchMultiCase( t *testing.T)  {
	for i:=0;i<5;i++{
		// switch case 支持多项
		// 默认不需要加break
		switch i {
		case 0,2:
			t.Log("Evem")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T)  {
	for i:=0;i<5;i++{
		// switch case condition表达
		switch  {
		case 0==i||2==i:
			t.Log("Evem")
		case 1==i||3==i:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}
