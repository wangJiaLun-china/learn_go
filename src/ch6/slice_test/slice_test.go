package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T)  {
	var s0 [] int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1:=[]int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	s2:=make([]int,3,5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0],s2[1],s2[2])
	s2= append(s2, 1)
	t.Log(s2[0],s2[1],s2[2],s2[3])
	s2= append(s2, 1)
	t.Log(s2[0],s2[1],s2[2],s2[3],s2[4])
	s2= append(s2, 1)
	t.Log(s2[0],s2[1],s2[2],s2[3],s2[4],s2[5])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T)  {
	s:=[]int{}
	for i := 0; i < 30; i++ {
		s = append(s,i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T)  {
	// 切片共享存储结构
	year:=[]string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer:=year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T)  {
	a:=[]int{1,2,3,4}
	b:=[]int{1,2,3,4}
	fmt.Println(a,b)
	// 切片不可以比较
	// slice can only be compared to nil
	//fmt.Println(a==b)
}