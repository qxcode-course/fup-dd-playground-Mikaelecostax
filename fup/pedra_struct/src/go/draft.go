package main

import
	"fmt"
func main() {
    var n int
    fmt.Scan(&n)
    vencedor:= -1
    MP := 999999
    for i := 0; i < n; i++ {
        var a, b int
        fmt.Scan(&a, &b)
        if a < 10 || b < 10 {
            continue
        }
    diferenca := a-b
    if diferenca < 0 {
        diferenca = -diferenca
    }
    if diferenca < MP {
        MP = diferenca
        vencedor = i
    }
 }
  if vencedor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(vencedor)
    }
}
