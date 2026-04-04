package main
import "fmt"
func main() {
    var a, soma, subtração int
    var b, multiplicação int
    var resto int
    var divisão float64
    fmt.Scan(&a, &b)
    soma = (a + b)
    subtração = (a - b) 
    multiplicação = (a * b)
    divisão = float64(a) / float64(b)
    resto = (a % b)
    fmt.Printf("%d\n", soma)
    fmt.Printf("%d\n", subtração)
    fmt.Printf("%d\n", multiplicação)
    fmt.Printf("%.2f\n", divisão)
    fmt.Printf("%d\n", resto)
}
