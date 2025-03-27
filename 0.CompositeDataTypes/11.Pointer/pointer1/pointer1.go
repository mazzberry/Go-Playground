package main

import "fmt"

func main() {
	i, j := 20, 40

	var ip *int = &i
	var jp *int = &j

	fmt.Println("address of i : i", &i)
	fmt.Println("address of i : i", ip)

	fmt.Println("address of j : j", &j)
	fmt.Println("address of j : j", jp)

	//

	i2 := i
	i2 = i2 + 2

	fmt.Println("value of i : ", i)
	fmt.Println("value of i2 : ", i2)

	fmt.Println("address of i : i", &i2)

	// to copy the value of i to ip2 using pointer
	// this is done by dereferencing the pointer *ip2
	// this way will change the value of i by the pointer

	var ip2 *int = &i

	fmt.Println("address of i : i", ip2) // the address of i2 is same as the address of i
	fmt.Println("value of ip2", *ip2)    // to get the value pointed by the pointer

	*ip2 = *ip2 + 6
	fmt.Println("value of i after incrementing ip2 : ", *ip2)
	fmt.Println("value of i : ", i)

}

