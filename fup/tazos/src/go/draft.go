package main
import "fmt"
func main() {
    var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	if n == 0 {
		fmt.Println("[ ]")
		return
	}
	maxFrequencia := 0
	frequenciaAtual := 1
	var maisRepetidos []int
	for i := 1; i <= n; i++ {
		if i < n && vetor[i] == vetor[i-1] {
			frequenciaAtual++
		} else {
			numeroAnterior := vetor[i-1]
			if frequenciaAtual > maxFrequencia {
				maxFrequencia = frequenciaAtual
				maisRepetidos = []int{numeroAnterior}
			} else if frequenciaAtual == maxFrequencia {
				maisRepetidos = append(maisRepetidos, numeroAnterior)
			}
			frequenciaAtual = 1
		}
	}
	fmt.Print("[")
	for _, v := range maisRepetidos {
		fmt.Printf(" %d", v)
	}
	fmt.Println(" ]")
}
