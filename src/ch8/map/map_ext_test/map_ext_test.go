package map_ext_test

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T)  {
	m:= map[int]func(op int)int{}
	m[1]= func(op int) int {return op}
	m[2]= func(op int) int {return op*op}
	m[3]= func(op int) int {return op*op*op}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T)  {
	// go 无内置set，基于map简单实现特性
	mySet:=map[int]bool{}
	n:=1
	mySet[n] = true
	if mySet[n] {
		t.Logf("%d is existing", n)
	}else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet))
	delete(mySet, n)
	if mySet[n] {
		t.Logf("%d is existing", n)
	}else {
		t.Logf("%d is not existing", n)
	}
}