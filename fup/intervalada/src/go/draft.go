package main
import "fmt"
func main() {
    var n, inferior, superior int
	fmt.Scan(&n, &inferior, &superior)
	contador := 0
	for i := 0; i < n; i++ {
		var numero int
		fmt.Scan(&numero)
		if numero >= inferior && numero <= superior {
			contador++
		}
	}
	fmt.Println(contador)
}
