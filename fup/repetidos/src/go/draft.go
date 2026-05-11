package main
import "fmt"
func main() {
    var p, n int
    var elemento int
    if _, err := fmt.Scan(&p, &n); err != nil {
		return
	}
	contador := 0
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&elemento); err == nil {
			if elemento == p {
				contador++
            }
		}
	}
	fmt.Println(contador)
}