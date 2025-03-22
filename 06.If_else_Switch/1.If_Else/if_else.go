package main

import "fmt"

func main() {

	// if statement only
	a := 10
	if a > 9 {
		println("a is greater than 9")
	}

	b := 4
	if b > 3 && b < 5 {
		fmt.Println(" b is whithing the range")
	}

	// if else statement
	c := 1
	d := 2

	if c > d {
		fmt.Println("c is greater than d")
	} else {
		fmt.Println("d is greater than c")

	}

	// else if statement
	age := 29
	if age < 18 {
		fmt.Println("You are minor")
	} else if age >= 18 && age <= 60 {
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are senior")
	}

	// nested if else statement
	/*
		if condition {
	  		//Do something
	  		if condition2 {
	    		//Do something
	  		} else {
	  //Do something
	}
	*/
	x := 10
	y := 20
	z := 30

	if x > y && a > z {
        fmt.Println("Biggest is x")
    } else if y > a && y > c {
        fmt.Println("Biggest is y")
    } else {
        fmt.Println("Biggest is z")
    }




}
