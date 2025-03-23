package main

func main() {

	a := []string{"a", "b"}
	checkAndPrint(a, 2)



}

func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}

}
