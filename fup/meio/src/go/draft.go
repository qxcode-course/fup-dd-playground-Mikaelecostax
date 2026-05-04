package main
import "fmt"
func main() {
    var a int
    var b int
    var c int
    var medio int
    fmt.Scan(&a, &b, &c)
    if (a > b && a < c) || (a < b && a > c) {
		medio = a
	} else if (b > a && b < c) || (b < a && b > c) {
		medio = b
	} else {
		medio = c
	}
    fmt.Println(medio)
}
