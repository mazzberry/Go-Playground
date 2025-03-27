package main

import (
	"encoding/json"
	"fmt"
)

const s string = "hello world!"

func main() {
	fmt.Println(s)

	// defining constant without needing declare type

	// named untype constant
	const a = 1
	const b = "circle"
	const c = 5.4
	const d = true
	const e = 'a'
	const f = 3 + 5i

	var u = 123      //Default hidden type is int
	var v = "circle" //Default hidden type is string
	var w = 5.6      //Default hidden type is float64
	var x = true     //Default hidden type is bool
	var y = 'a'      //Default hidden type is rune
	var z = 3 + 5i   //Default hidden type is complex128

	fmt.Println("")
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", d, d)
	fmt.Printf("Type: %T Value: %v\n", e, e)
	fmt.Printf("Type: %T Value: %v\n", f, f)

	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// V-Course

	fmt.Print(Pending, Shipped, Cancelled, Returned, Completed, Delivered)

	order1 := Order{ID: 1, Price: 100, status: Pending}
	order2 := Order{ID: 2, Price: 200, status: Shipped}
	order3 := Order{ID: 3, Price: 300, status: Cancelled}

	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println(order3)

	order1Json, _ := json.Marshal(order1)
	order2Json, _ := json.Marshal(order2)
	order3Json, _ := json.Marshal(order3)
	

	fmt.Println(string(order1Json))
	fmt.Println(string(order2Json))
	fmt.Println(string(order3Json))

}

type Order struct {
	ID     int
	Price  int
	status Orderstatus
}

type Orderstatus int

const (
	Pending   Orderstatus = iota // 0
	Shipped                      //1
	Cancelled                    //2
	Returned                     //3
	Completed                    //4
	Delivered                    //5
)
