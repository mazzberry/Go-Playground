package main

import "fmt"

type Number interface{
	int | int16 | int32 | int64 | float32 | float64
}


func main() {
	x := Sum(2, 7)
	fmt.Printf("%d\n", x)

	y := Sum(4.3, 2.2)
	fmt.Printf("%.2f\n", y)
}

// func SumInt(a, b int) int {
// 	return a + b
// }

// func SumFloat(a, b float64) float64 {
// 	return a + b
// }


func Sum[T Number] (a, b T) T{
	return a + b
}


// \\\\\\\\\\\

// func Call [Input any, Output any](headers string, uri string, input Input){
// 	// http request(header uri, input)
// 	//

// }