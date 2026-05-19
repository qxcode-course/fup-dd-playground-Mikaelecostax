package main
import "fmt"
func main() {
    var n int
	fmt.Scan(&n)
	blocos := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&blocos[i])
	}
	parkour := 0
	for i := 0; i < n-1; i++ {
		diff := blocos[i] - blocos[i+1]
		if diff < 1 {
			diff = -diff
		}
		if diff > 0 {
			parkour++
		}
	}
	fmt.Println(parkour)
}
