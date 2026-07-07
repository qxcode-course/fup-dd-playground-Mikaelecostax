package main
import "fmt"
func main() {
    var n int 
    fmt.Scan(&n)
    matriz:= make([][]int, n)
    for i := 0; i < n; i++ {
        matriz[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }
    MS := -1
    ColV := 0
    for j := 0; j < n; j++ {
        somac := 0
        for i := 0; i < n; i++ {
            elemento := matriz[i][j]
            somac += elemento * elemento
        }
        if somac > MS{
            MS = somac
            ColV = j 
        }
    }
    fmt.Println(ColV)
}