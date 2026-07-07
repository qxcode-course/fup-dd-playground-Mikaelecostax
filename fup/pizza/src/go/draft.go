package main
import "fmt"
type restaurante struct {
    nome string
    pontos int
}
func encontrarMP(restaurantes []restaurante) string {
    vencedor := restaurantes[0]
    for i := 1; i < len(restaurantes); i++ {
        atual := restaurantes[i]
        if atual.pontos > vencedor.pontos {
            vencedor = atual
        } else if atual.pontos == vencedor.pontos {
            if atual.nome < vencedor.nome {
                vencedor = atual
            }
        }
    }
    return vencedor.nome
}
func main () {
    var n int
    fmt.Scan(&n)
    restaurantes := make([]restaurante, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&restaurantes[i].nome, &restaurantes[i].pontos)
    }
    vencedornome := encontrarMP(restaurantes)
    fmt.Println(vencedornome)
}