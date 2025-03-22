package main  
  
import "fmt"  
  
func main() {
	const (
		A = iota
		B
		C
		D = B + C
		E
		F
		G = iota
		H
		I = H
		J
		K
	)
	fmt.Println(A, B, C, D, E, F, G, H, I, J, K)
}