package main

import "fmt"

func main() {
	var companyProfile = map[string]string{
		"name":    "companyName",
		"address": "sampleAddress",
	}

	var editorMap = map[string]string{}

	for key, value := range companyProfile {
		editorMap[key] = value
	}

	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])
	fmt.Println(editorMap["name"], "\t", editorMap["address"])

	editorMap["name"] = "new address"
	editorMap["address"] = "new address"

	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])
	fmt.Println(editorMap["name"], "\t", editorMap["address"])
}
