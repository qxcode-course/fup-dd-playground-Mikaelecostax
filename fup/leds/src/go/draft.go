package main
import (
	"fmt"
)
func main() {
	ledsPorDigito := map[rune]int{
		'1': 2,
		'7': 3,
		'4': 4,
		'2': 5, '3': 5, '5': 5,
		'0': 6, '6': 6, '9': 6,
		'8': 7,
	}
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	for i := 0; i < n; i++ {
		var v string
		fmt.Scan(&v)
		totalLeds := 0
		for _, digito := range v {
			totalLeds += ledsPorDigito[digito]
		}
		fmt.Printf("%d leds\n", totalLeds)
	}
}
