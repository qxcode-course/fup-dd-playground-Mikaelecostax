package main
import( "fmt"
    "math"
)
func main() {
   var n int
	fmt.Scan(&n)
	machos := make(map[int]int)
	femeas := make(map[int]int)
	for i := 0; i < n; i++ {
		var animal int
		fmt.Scan(&animal)
		if animal > 0 {
			machos[animal]++
		} else if animal < 0 {
			especieFemea := int(math.Abs(float64(animal)))
			femeas[especieFemea]++
		}
	}
	totalCasais := 0
	for especie, qtdMachos := range machos {
		qtdFemeas := femeas[especie]
		if qtdMachos < qtdFemeas {
			totalCasais += qtdMachos
		} else {
			totalCasais += qtdFemeas
		}
	}
	fmt.Println(totalCasais)
}
