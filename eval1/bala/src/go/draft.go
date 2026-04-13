package main
import ("fmt" 
 "math"
)
func main() {
    var x1, x2 float64
    var y1, y2 float64
    var distanciaab float64
    var pontox, pontoy float64
    var quadrado float64
    var quadrado2 float64
    fmt.Scan(&x1, &x2, &y1, &y2, &quadrado, &quadrado2)
    pontox = (x2 - x1)
    quadrado= math.Pow(pontox, 2)
    pontoy = (y2 - y1)
    quadrado2=math.Pow(pontoy, 2)
    distanciaab = math.Sqrt(quadrado + quadrado2)
    fmt.Printf("%.2f\n", distanciaab)
}
