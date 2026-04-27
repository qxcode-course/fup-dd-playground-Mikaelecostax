package main
import "fmt"
func main() {
    var v1, v2, v3 int
    fmt.Scan(&v1, &v2, &v3)
    if v1 < v2+v3 && v2 < v1+v3 && v3 < v1+v2{
        fmt.Println("True")
    } else {
        fmt.Println("False")
    } 
}
