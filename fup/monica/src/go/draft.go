package main
import "fmt"
func main() {
    var m int
    var a int 
    var b int 
    fmt.Scan(&m, &a, &b)
    c := m - (a + b)
    maisvelho := a 
    if b > maisvelho {
		maisvelho = b
	}
	if c > maisvelho {
		maisvelho = c
	}
    fmt.Println(maisvelho)
}
