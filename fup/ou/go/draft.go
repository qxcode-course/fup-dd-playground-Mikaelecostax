package main
import  ( "fmt"
       "math"
)
func main() {
    var a, maior float64
    fmt.Scan(&a)
    maior = math.Max(a)
    fmt.Println("%.1f\n")
}
