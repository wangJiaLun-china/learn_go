package array_test

import "testing"

func TestArrayInit(t *testing.T)  {
	var arr[3]int
	arr1 := [4] int{1,2,3,4}
	arr2 := [...] int{1,2,3,4}
	arr1[1] = 7
	t.Log(arr[1], arr[2])
	t.Log(arr1[1], arr1[2])
	t.Log(arr2[1], arr2[2])
}

func TestArrayTravel(t *testing.T)  {
	arr2 := [...] int{2,3,3,4}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}
	for i := range arr2 {
		t.Log(i)
	}
	for index, item := range arr2 {
		t.Log(index,item)
	}
}

func TestArraySection(t *testing.T)  {
	a:= [...]int {1,2,3,4,5}
	// 不支持负数下标
	t.Log(a[1:2])
	t.Log(a[1:3])
	t.Log(a[1:len(a)])
	t.Log(a[1:])
	t.Log(a[:2])
}