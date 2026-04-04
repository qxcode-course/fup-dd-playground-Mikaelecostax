package main
import "fmt"
func main() {
    var numero int
    fmt.Scan(&numero)
    if numero % 7 == 0 {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
