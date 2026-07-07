package main
import "fmt"
func main() {
    var quants int 
    fmt.Scan(&quants)
    alturas := make([]float64, quants)
    soma := 0.0
    for i := 0; i < quants; i++ {
        fmt.Scan(&alturas[i])
        soma += alturas[i]
    }
    media := soma / float64(quants)
    fmt.Printf("%.2f\n", media)
    for i := 0; i < quants; i++ {
        if alturas[i] < media {
           fmt.Print("P")
        } else if alturas[i] > media {
           fmt.Print("G")
        } else {
           fmt.Print("M")
        }
        if i < quants-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println( )
}