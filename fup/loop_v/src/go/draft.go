package main
import "fmt"
func main() {
    var a int
    var b int 
    fmt.Scan(&a, &b)
    fmt.Printf("[ ")
    for   { 
        if a >= b {
            break
        }
        if a % 2 == 0 {
            a++ 
            continue
        }
        fmt.Printf("%d ", a)
        a++
    }
    fmt.Println("]")
}
