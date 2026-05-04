package main
import ("fmt"
"math" )
func main() {
    var c int
    var a int
    fmt.Scan(&c, &a)
    capacidadealunos := c - 1
    viagens := float64(a) / float64(capacidadealunos)
	resultado := int(math.Ceil(viagens))
    fmt.Println(resultado)
}
