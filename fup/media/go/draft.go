package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    media:= float64(a + b) / 2
    fmt.Printf("%.1f\n",media)
}
