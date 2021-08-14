package main

import (
	"math"
	"syscall"
)

func toDoublePtr(x float64) uintptr {
	var a uint64 = math.Float64bits(x)
	return uintptr(a)
} //参见https://bbs.csdn.net/topics/396300997?list=68399813 除int,void外所有传参都需要转换
func norm2(a float64, b float64, m int) (ret0, ret1 float64) {
	arr1 := [3]string{"norm2", "norm2_int", "norm2_fl"}
	dll1 := syscall.NewLazyDLL("..\\cpp\\Debug\\DLLTest1.dll")
	n2 := dll1.NewProc(arr1[m])
	if n2 == nil {
		println("Unable to find required process")
		return 0, 0
	} else {
		r, r1, _ := n2.Call(toDoublePtr(a), toDoublePtr(b)) // double型函数的返回值竟然在r1
		ret0 := math.Float64frombits(uint64(r))             //采用整型返回值时r和r1都能返回正确答案,但采用浮点型返回值时只有r1正确,r必定为0
		ret1 := math.Float64frombits(uint64(r1))
		return ret0, ret1
	}
}
