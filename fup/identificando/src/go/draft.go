package main
import ("fmt"
    "strings" )
func main() {
    var elemento string
    var tipo string
    primeiro := true 
    for {
    _, err := fmt.Scan(&elemento)
    if err != nil {
        break
    }
    if strings.ContainsAny(elemento, "abcdefghijklmnopqrstuvwxyz") {
        tipo = "str"
    } else if strings.Contains(elemento, ".") {
        tipo = "float"
    } else {
        tipo = "int"
    }
    if primeiro {
        fmt.Print(tipo)
        primeiro = false
    } else {
        fmt.Print(" " + tipo)
    }
    }
    fmt.Println()
}