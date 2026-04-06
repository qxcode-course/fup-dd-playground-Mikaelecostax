package main
import "fmt"
func main() {
    var a int
    var b int
    var c string
    var resultado int
    fmt.Scan(&a, &b, &c)
    if c == "+" {
        resultado = a + b
        fmt.Printf("%d\n", resultado)
    } else if c == "-"{
        resultado = a - b 
        fmt.Printf("%d\n", resultado)
    } else if c == "*" {
        resultado = a*b
        fmt.Printf("%d\n", resultado)
    } else if c == "/" {
        resultado = a/b
        fmt.Printf("%d\n", resultado)
    }
   
}
