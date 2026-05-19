package main
import "fmt"
func main() {
    var n int
	fmt.Scan(&n)
	var alunos []int  
	var servidores []int 
	for i := 0; i < n; i++ {
		var pessoa int
		fmt.Scan(&pessoa)
		if pessoa%2 != 0 {
			alunos = append(alunos, pessoa)
		} else {
			servidores = append(servidores, pessoa)
		}
	}
	imprimirFila(servidores)
}
func imprimirFila(fila []int) {
	fmt.Print("[")
	for i, v := range fila {
		fmt.Print(v)
		if i < len(fila)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}
