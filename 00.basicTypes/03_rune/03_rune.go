package main

import "fmt"

func stringAndRuneCompare(){
	s := "helloبه"
	fmt.Printf("%s type: %T, len: %d\n", s, s, len(s))

	rs := []rune(s)
	fmt.Printf("%s type: %T, len: %d\n", rs, rs, len(rs))
}

/* []rune, byte and string could convert to each other */

func main() {
	var r rune = 'k'
	fmt.Printf("%c %T %U", r, r, r)

	stringAndRuneCompare()

}