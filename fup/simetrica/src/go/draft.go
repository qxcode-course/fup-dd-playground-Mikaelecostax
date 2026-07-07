package main
import "fmt"
func main() {
    var matriz [3][3]int 
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }
    simetrica := true
    if matriz[0][1] != matriz[1][0] ||
    matriz[0][2] != matriz[2][0] ||
    matriz[1][2] != matriz[2][1] {
        simetrica = false
    }
    if simetrica {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}