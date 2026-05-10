package main
import "fmt"
func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
		return
	}
    if n == 0 {
		fmt.Println() 
		return
	}
    if n <= 0 {
		return
	}
    vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println(vetor[i])
    }
}
