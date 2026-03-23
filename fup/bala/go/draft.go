package main
import ("fmt"
       "math"
)
func main() {
    var x1, y1, distancia, distanciax,  distanciay float64
    var x2, y2, distanciat float64
    fmt.Scan(&x1, &y1, &x2, &y2)
    distanciax = (x2-x1)
    distanciax= math.Pow(distanciax, 2) 
    distanciay = (y2-y1)
    distanciay= math.Pow(distanciay, 2)
    distanciat = distanciax + distanciay
    distancia= math.Sqrt(distanciat)
    fmt.Printf("%.2f\n", distancia)
}
