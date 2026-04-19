package main
import "fmt"
func main() {
    var p, d1, d2, soma int
    var resultadopar bool
    fmt.Scan(&p, &d1, &d2, &resultadopar, &soma)
    soma = d1 + d2
    resultadopar = (soma % 2 == 0)
    if p == 0 {
        fmt.Println(resultadopar)
    } else if p == 1 {
        fmt.Println(resultadopar)
    }
}