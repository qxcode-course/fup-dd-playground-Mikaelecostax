package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var texto string
	if scanner.Scan() {
		texto = scanner.Text()
	}
	var alvo string
	if scanner.Scan() {
		alvo = scanner.Text()
	}
	var nova string
	if scanner.Scan() {
		nova = scanner.Text()
	}
	lenTexto := len(texto)
	lenAlvo := len(alvo)
	if lenAlvo == 0 {
		fmt.Println(texto)
		return
	}
	var resultado []byte
	i := 0
	for i < lenTexto {
		if i+lenAlvo <= lenTexto && texto[i:i+lenAlvo] == alvo {
			resultado = append(resultado, nova...)
			i += lenAlvo
		} else {
			resultado = append(resultado, texto[i])
			i++
		}
	}
	fmt.Println(string(resultado))
}
