package main
import "fmt"
func main() {
    var nl, nc int
    fmt.Scan(&nl, &nc)
    matriz := make([][]int, nl)
    for i := 0; i < nl; i++ {
        matriz[i] = make([]int, nc)
        for j := 0; j < nc; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }
    totalocor := 0
    for j := 0; j < nc; j++ {
        for i := 0; i < nl-1; i++ {
            if matriz[i][j] > matriz[i+1][j]{
                totalocor++
            }
        }
    }
        fmt.Println(totalocor)
}