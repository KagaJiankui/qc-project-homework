package main

import (
	"fmt"
	"math"
	"syscall"
)

func main() {
	// fmt.Println("Hello, 世界")
	var result float64 = norm2(1.36, 2.5)
	fmt.Println("Result: ", result)
}
func toDoublePtr(x float64) uintptr {
	var a uint64 = math.Float64bits(x)
	return uintptr(a)
} //参见https://bbs.csdn.net/topics/396300997?list=68399813 除int,void外所有传参都需要转换
func norm2(a, b float64) float64 {
	dll1 := syscall.NewLazyDLL("..\\cpp\\Debug\\DLLTest1.dll")
	n2 := dll1.NewProc("norm2")
	if n2 == nil {
		println("Unable to find required process")
		return 0
	} else {
		_, r1, _ := n2.Call(toDoublePtr(a), toDoublePtr(b)) // double型函数的返回值竟然在r1
		rp := math.Float64frombits(uint64(r1))
		return rp
	}

}
