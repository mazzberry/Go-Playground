package main

import "fmt"

func main() {

	var companyProfile = map[string]string{
		"name":    "companyName",
		"address": "sampleAddress",
	}

	var editorMap = companyProfile

	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])

	fmt.Println(editorMap["name"], "\t", editorMap["address"])

	editorMap["name"] = "new name"
	editorMap["address"] = "new address"

	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])

	fmt.Println(editorMap["name"], "\t", editorMap["address"])

}
