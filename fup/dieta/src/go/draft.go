package main
import "fmt"
func main() {
    var n int
    var caloriasDia int
    fmt.Scan(&n)
    soma := 0
    for i := 0; i < n; i++ {
        fmt.Scan(&caloriasDia)
        soma += caloriasDia
    }
    media := float64(soma) / float64(n)
    fmt.Printf("%.1f\n", media)
}
