package main
import "fmt"
func main() {
    var n int
	fmt.Scan(&n)
	var unicos []int
	for i := 0; i < n; i++ {
		var animal int
		fmt.Scan(&animal)
		existe := false
		for _, u := range unicos {
			if u == animal {
				existe = true
				break
			}
		}
		if !existe {
			unicos = append(unicos, animal)
		}
	}
	tam := len(unicos)
	for i := 0; i < tam-1; i++ {
		for j := 0; j < tam-i-1; j++ {
			if unicos[j] > unicos[j+1] {
				unicos[j], unicos[j+1] = unicos[j+1], unicos[j]
			}
		}
	}
	fmt.Print("")
	for i, v := range unicos {
		fmt.Print(v)
		if i < len(unicos)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}
