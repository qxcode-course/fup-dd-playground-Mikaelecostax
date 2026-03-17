package main
import "fmt"
func main() {
    var a, b float64
    fmt.Scan(&a, &b)
    b = (1.8 * a) + 32
 fmt.Printf("%.6f\n", b)
}
