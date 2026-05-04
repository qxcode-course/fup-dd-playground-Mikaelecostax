package main
import "fmt"
func main() {
    var n, d, a int
    var cliques int
    fmt.Scan(&n, &d, &a)
    if d >= a {
		cliques = d - a
	} else {
		cliques = (n - a) + d
	}
    fmt.Println(cliques)
}
