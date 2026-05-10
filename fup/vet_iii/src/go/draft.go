package main
import "fmt"
func main() {
    var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	if n <= 0 {
		fmt.Println("[]")
		return
	}
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	fmt.Print("[")
	for i := 0; i < n; i++ {
		fmt.Print(vetor[i])
		if i < n-1 {
			fmt.Print(", ")
		}
    }
	fmt.Println("]")
}
