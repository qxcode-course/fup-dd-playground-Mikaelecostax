package main
import ( "fmt"
"math" )
func main() {
    var c, bananas, goiabas, mangas int
    fmt.Scan(&c, &bananas, &goiabas, &mangas)
    totalfrutas := bananas + goiabas + mangas
    minutos := math.Ceil(float64(totalfrutas) / float64(c))
    fmt.Printf("%.0f\n", minutos)
}
