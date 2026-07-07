package main
import "fmt"
func main() {
    var linhas, colunas int
    fmt.Scan(&linhas, &colunas)
    matrizA:= make([][]int, linhas)
    for i := 0; i < linhas; i++ {
        matrizA[i] = make([]int, colunas)
        for j := 0; j < colunas; j++ {
            fmt.Scan(&matrizA[i][j])
        }
    }
    matrizB := make([][]int,linhas)
    for i := 0; i < linhas; i++ {
        matrizB[i] = make([]int, colunas)
        for j := 0; j < colunas; j++ {
            fmt.Scan(&matrizB[i][j])
        }
    }
    for i := 0; i < linhas; i++ {
        fmt.Print("[ ")
        for j := 0; j < colunas; j++ {
            soma:= matrizA[i][j] + matrizB[i][j]
            if j == colunas-1 {
                fmt.Print(soma)
            } else {
                fmt.Print(soma, " ")
            }
        }
        fmt.Println(" ]")
    }
}