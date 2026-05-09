package main
import "fmt"
func main() {
    var a int
    var b int
    _, err := fmt.Scan(&a, &b)
    soma := 0
    if err != nil {
		return
	}
    if a > b {
		fmt.Println("invalido")
		return
	}
    for i := a; i <= b; i++ {
		if i % 2 == 0 {
			soma += i
		}
    }
    fmt.Println(soma)
}
