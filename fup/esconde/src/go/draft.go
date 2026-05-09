package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
    for i := n; i >= 0; i-- {
		if i%2 == 0 {
			fmt.Println(i)
		}
    }
}
