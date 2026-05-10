package main
import "fmt"
func main() {
    var chuteChico, chuteCebolinha, qtdAnimais int
    var tipo string
	fmt.Scan(&chuteChico)
	fmt.Scan(&chuteCebolinha)
	fmt.Scan(&qtdAnimais)
    totalPatas := 0
	for i := 0; i < qtdAnimais; i++ {
		fmt.Scan(&tipo)
		switch tipo {
		case "v": 
			totalPatas += 4
		case "g": 
			totalPatas += 2
		case "c":
			totalPatas += 4
		}
	}
	difChico := chuteChico - totalPatas
	if difChico < 0 {
		difChico = -difChico
	}
    difCebolinha := chuteCebolinha - totalPatas
	if difCebolinha < 0 {
		difCebolinha = -difCebolinha
	}
	fmt.Println(totalPatas)

	if difChico < difCebolinha {
		fmt.Println("Chico Bento")
	} else if difCebolinha < difChico {
		fmt.Println("Cebolinha")
	} else {
		fmt.Println("empate")
	}
}
