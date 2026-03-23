package main
import ("fmt"
 "math"
)
func main() {
    var a float64
    var b float64
    var resultado float64
    fmt.Scan(&a, &b)
    resultado = math.Abs(a-b)
    

    fmt.Println( resultado)
}
