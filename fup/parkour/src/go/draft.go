package main
import  "fmt"
func main() {
    var n int
	fmt.Scan(&n)
	blocos := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&blocos[i])
	}
	parkour := 0
	for i := 0; i < n-1; i++ {
		diferenca := blocos[i] - blocos[i+1]
		if diferenca < 0 {
			diferenca = -diferenca
		}
		if diferenca > 1 {
			parkour++
		}
	}
	fmt.Println(parkour)
}