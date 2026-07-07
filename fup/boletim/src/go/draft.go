package main
import "fmt"
func main() {
    var matriz[2][3]int 
    somaT := 0
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&matriz[i][j])
            somaT += matriz[i][j]
        }
    }
    fmt.Println(somaT)
}