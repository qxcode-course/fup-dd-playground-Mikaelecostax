package main
import ("fmt" 
"math"
)
func main() {
    var a float64
    var b float64
    var maior float64
    fmt.Scan(&a, &b)
maior = math.Max(a,b)
fmt.Println(maior)
}
