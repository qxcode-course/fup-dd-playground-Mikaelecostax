package main
import "fmt"
func main() {
    var a, b, c, h, l int
    fmt.Scan(&a, &b, &c, &h, &l)
    passa := false
    if (a <= h && b <= l) || (a <= l && b <= h) {
		passa = true
	} else if (a <= h && c <= l) || (a <= l && c <= h) { 
		passa = true
	} else if (b <= h && c <= l) || (b <= l && c <= h) {
		passa = true
	}
    if passa {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
