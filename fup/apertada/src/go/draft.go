package main
import "fmt"
func main() {
    var atual, menor int
    fmt.Scan(&atual)
    menor = atual
    for i := 0; i < 4; i++ {
        fmt.Scan(&atual)
        if atual < menor {
            menor = atual
        }
    }
    fmt.Println(menor)
}
