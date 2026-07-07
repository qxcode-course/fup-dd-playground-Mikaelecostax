package main
import "fmt"
func main() {
    var matriz [3][3]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }
    somaref := matriz[0][0] + matriz[0][1] + matriz[0][2]
    emagico := true
    for i := 1; i < 3; i++ {
        somal := matriz[i][0] + matriz[i][1] + matriz[i][2]
        if somal != somaref {
            emagico = false
        }
    }
    for j := 0; j < 3; j++ {
        somacol := matriz[0][j] + matriz[1][j] + matriz[2][j]
        if somacol != somaref {
            emagico = false
        }
    }
    somadiagp := matriz[0][0] + matriz[1][1] + matriz[2][2]
    if somadiagp != somaref {
        emagico = false
    }
    somadiags := matriz[0][2] + matriz[1][1] + matriz[2][0]
    if somadiags != somaref{
        emagico = false 
    }
    if emagico {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}