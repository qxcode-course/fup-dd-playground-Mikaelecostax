package main
import "fmt"
func main() {
    var a int
    var b int
    fmt.Scan(&a, &b)
    fmt.Printf("[ ")
    if a == b {
        fmt.Println()
        return
    }
    step := 1
    if a > b {
        step = -1 
    }
    for i := a; i!=b; i+= step {
        fmt.Printf("%d ",i)
    }
    fmt.Println("]")
}
