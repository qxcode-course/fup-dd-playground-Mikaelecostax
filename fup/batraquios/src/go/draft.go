package main
import "fmt"
func main() {
    var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	vetor1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor1[i])
	}
    var m int
	if _, err := fmt.Scan(&m); err != nil {
		return
	}
	vetor2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&vetor2[i])
	}
    todosContidos := true
    for _, elem1 := range vetor1 {
		encontrou := false
		for _, elem2 := range vetor2 {
			if elem1 == elem2 {
				encontrou = true
				break
			}
		}
        if !encontrou {
			todosContidos = false
			break
		}
	}
    if todosContidos {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}
