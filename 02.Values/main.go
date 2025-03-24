package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {
	var r rune = 'k'

	fmt.Printf("%c %T %U", r, r, r)
	var a = bits.UintSize
	fmt.Println(a)

	var fnum int
	var fnum8 int8
	var fnum16 int16
	var fnum32 int32
	var fnum64 int64

	

	fmt.Printf("num %d bytes\n", unsafe.Sizeof(fnum)) 
	fmt.Printf("num8 %d bytes\n", unsafe.Sizeof(fnum8))
	fmt.Printf("num16 %d bytes\n", unsafe.Sizeof(fnum16))
	fmt.Printf("num32 %d bytes\n", unsafe.Sizeof(fnum32)) 
	fmt.Printf("num64 %d bytes\n", unsafe.Sizeof(fnum64))


	var floatnum32 float32 
	var floatnum64 float64 

	fmt.Printf("floatnum32 %d bytes\n", unsafe.Sizeof(floatnum32))
	fmt.Printf("floatnum64 %d bytes\n", unsafe.Sizeof(floatnum64))
	

	

}
