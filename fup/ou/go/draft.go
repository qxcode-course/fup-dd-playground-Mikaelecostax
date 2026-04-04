package main
import   "fmt"

func main() {
    var numero int
    fmt.Scan(&numero)
    if numero == 3 || numero == 5 {
        fmt.Println("SIM")
        } else {
            fmt.Println("NAO")
        }
}
