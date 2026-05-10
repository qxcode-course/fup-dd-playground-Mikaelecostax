package main
import ( "fmt"
    "strings" )
func main() {
    var a, b int
    var resultado []string
	fmt.Scan(&a, &b)
	inicio := a
	fim := b
    for i, j := inicio, fim; i <= fim; i, j = i+1, j-1 {
		resultado = append(resultado, fmt.Sprintf("%d", i))
		resultado = append(resultado, fmt.Sprintf("%d", j))
    }
    fmt.Printf("[ %s ]\n", strings.Join(resultado, " "))
}
