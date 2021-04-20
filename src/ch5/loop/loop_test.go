package loop_test

import "testing"

func TestWhileLoop(t *testing.T)  {
	n:= 0
	/* while(n<5)*/
	for n < 5 {
		t.Log(n)
		n++
	}
	// 无限循环
	for{
		n++
		if n == 10 {
			break
		}
	}
	t.Log(n)
}
