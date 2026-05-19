package main
import "fmt"
func main() {
    var n int
	fmt.Scan(&n)
	botasD := make([]int, 61)
	botasE := make([]int, 61)
	for i := 0; i < n; i++ {
		var tamanho int
		var pe string
		fmt.Scan(&tamanho, &pe)
		if pe == "D" {
			botasD[tamanho]++
		} else if pe == "E" {
			botasE[tamanho]++
		}
	}
	totalPares := 0
	for i := 30; i <= 60; i++ {
		if botasD[i] < botasE[i] {
			totalPares += botasD[i]
		} else {
			totalPares += botasE[i]
		}
	}
	fmt.Println(totalPares)
}
