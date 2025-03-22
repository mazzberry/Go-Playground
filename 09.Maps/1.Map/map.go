package main

import "fmt"

func main() {

	myMap := make(map[int]string)
	myKey := 13
	myMap[myKey] = "thirteen"

	fmt.Println(myMap)
	fmt.Println(myMap[myKey])

	/*

		// Using Keyword var
		var sampleMap = map[keyType]valueType{keyName1:value1, keyName2:value2, ...}
		var sampleMap map[keyType]valueType = map[keyType]valueType{}

		// =: short variable declaration
		sampleMap := map[keyType]valueType{keyName1:value1, keyName2:value2, ...}

		// Using <make> function
		var sampleMap = make(map[keyType]valueType)
		sampleMap := make(map[keyType]valueType)

	*/
}
