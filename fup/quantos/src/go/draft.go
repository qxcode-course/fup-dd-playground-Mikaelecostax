package main
import "fmt"
func main() {
    var a int
    var b int
    var c int
    fmt.Scan(&a, &b, &c)
    if a == b && a == c  { 
        fmt.Print("3\n")
    } else if a == b && a != c {
        fmt.Print("2\n")
    } else if a != b && a == c {
        fmt.Print("2\n")
    } else if b == c && b != a {
        fmt.Print("2\n")
    } else if b != c && b == a {
        fmt.Print("2\n")
    } else if c != a && c != b {
        fmt.Print("0\n")
    }
  
}
