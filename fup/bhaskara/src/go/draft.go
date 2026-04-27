package main
import ("fmt"
"math"
)
func main() {
    var a, b, c, delta, x, x1, x2 float64
    fmt.Scan(&a,&b,&c, &delta, &x, &x1, &x2)
    delta = math.Pow(b, 2) - (4 * a * c)
    if delta < 0 {
        fmt.Println("nao ha raiz real")
    } else if delta == 0 {
        x = -b / (2 * a)
        fmt.Printf("%.2f\n", x)
    } else {
        raizDelta := math.Sqrt(delta)
		x1 = (-b + raizDelta) / (2 * a)
		x2 = (-b - raizDelta) / (2 * a)
        fmt.Printf("%.2f\n", x1)
		fmt.Printf("%.2f\n", x2)
    }
   
}
