package main

import (
	"fmt"
)


func main(){
	var sampleMap = map[string]bool{}
	var otherSample = make(map[string]uint8)
	var nilMap map[bool]bool

	sampleMap["Condition1"] = true
	sampleMap["Condition2"] = false

	otherSample["foo"] = 1

	fmt.Println(len(sampleMap))    // 2
	fmt.Println(len(otherSample))  // 1
	fmt.Println(len(nilMap))       // 0 (len nil is zero)

}