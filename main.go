package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	// fmt.Println("Hello, 世界")
	var result float64
	result = norm2(3, 4)
	fmt.Println("Result: ", result)
}
func toPtr(x float64) uintptr {
	return uintptr(unsafe.Pointer(&x))
}
func norm2(a, b float64) float64 {
	dll1 := syscall.NewLazyDLL("cpp\\Debug\\DLLTest1.dll")
	println("call DLL: ", dll1.Name)
	n2 := dll1.NewProc("norm2")
	// println(n2.Addr())
	ret, _, _ := n2.Call(toPtr(a), toPtr(b))
	rp := (*float64)(unsafe.Pointer(ret))
	return float64(*rp)
}
