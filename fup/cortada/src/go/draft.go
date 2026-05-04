package main
import "fmt"
func main() {
    var b int
    var t int
    fmt.Scan(&b, &t)
    const altura = 70
	const areaTotalMetade = (160 * 70) / 2
    areaFelix := ((b + t) * altura) / 2
    if areaFelix > areaTotalMetade {
		fmt.Println(1)
	} else if areaFelix < areaTotalMetade {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}

