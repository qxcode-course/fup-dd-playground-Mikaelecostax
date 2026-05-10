package main
import "fmt"
func main() {
    var p, s, e int
    fmt.Scan(&p, &s, &e)
    atual := 0
	for {
		salto := atual + s
		if salto >= p {
			fmt.Printf("%d saiu\n", atual)
			break
		}
		fmt.Printf("%d %d\n", atual, salto)
		atual = salto - e
	}
}


