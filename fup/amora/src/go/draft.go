package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var frase string
	if scanner.Scan() {
		frase = scanner.Text()
	}
	var trecho string
	if scanner.Scan() {
		trecho = scanner.Text()
	}
	if len(trecho) == 0 || len(trecho) > len(frase) {
		fmt.Println(0)
		return
  }
	contador := 0
	tamanhoTrecho := len(trecho)
	for i := 0; i <= len(frase)-tamanhoTrecho; i++ {
		pedaco := frase[i : i+tamanhoTrecho]
		if pedaco == trecho {
			contador++
		}
	}
	fmt.Println(contador)
}