package main
import "fmt"
func main() {
    var p, d1, d2, soma int
    fmt.Scan(&p, &d1, &d2)
    soma = d1 + d2
    if soma % 2 == 0  {
        fmt.Println(p)
    } else {
        fmt.Println(1-p)
    } 
   }