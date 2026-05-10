package main
import "fmt"
func main() {
    var h, p, f, d int
    fmt.Scan(&h, &p, &f, &d)
    atual := f
    for {
		atual += d

		if atual < 0 {
			atual = 15
		} else if atual > 15 {
			atual = 0
		}
		if atual == h {
			fmt.Println("S")
			break
		}
		if atual == p {
			fmt.Println("N")
			break
		}
    }
}
