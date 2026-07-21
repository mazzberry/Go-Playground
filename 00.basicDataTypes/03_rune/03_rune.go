package main

import "fmt"

func stringAndRuneCompare(){
	s := "helloبه"
	fmt.Printf("%s type: %T, len: %d\n \n", s, s, len(s))

	fmt.Println("-----")

	rs := []rune(s)
	fmt.Printf("%s type: %T, len: %d\n \n", string(rs), rs, len(rs))

	fmt.Println("-----")
}

/* []rune, byte and string could convert to each other */

func main() {
	var r rune = 'k'
	fmt.Printf("%c %T %U \n \n", r, r, r)

	fmt.Println("-----")

	stringAndRuneCompare()

}