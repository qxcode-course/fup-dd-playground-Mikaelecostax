package main
import "fmt"
func main() {
    var n, x, y, s int
	var c string
    fmt.Scan(&n, &x, &y, &c, &s)
    switch c {
	case "U":
		y = (y - s) % n
	case "D": 
		y = (y + s) % n
	case "L": 
		x = (x - s) % n
	case "R": 
		x = (x + s) % n
	} 
    if x < 0 {
		x += n
	}
	if y < 0 {
		y += n
	}
    fmt.Printf("%d %d\n", x, y)
}
