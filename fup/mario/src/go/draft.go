package main
import (
	"fmt"
)
func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	alturas := make([]int, n)
	maiorAltura := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&alturas[i])
		if alturas[i] > maiorAltura {
			maiorAltura = alturas[i]
		}
	}
	for nivel := maiorAltura; nivel >= 1; nivel-- {
		for col := 0; col < n; col++ {
			if alturas[col] >= nivel {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}
