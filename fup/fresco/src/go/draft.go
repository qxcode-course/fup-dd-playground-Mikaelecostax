package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func ehVogal(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		palavras := strings.Fields(texto)
		if len(palavras) == 0 {
			fmt.Println("")
			return
		}
		resultado := palavras[0]
		for i := 1; i < len(palavras); i++ {
			proxima := palavras[i]
			ultimoCharRes := resultado[len(resultado)-1]
			primeiroCharProx := proxima[0]
			if ehVogal(ultimoCharRes) && ehVogal(primeiroCharProx) {
				if ultimoCharRes == primeiroCharProx {
					resultado += proxima[1:]
				} else {
					novaJuncao := resultado + proxima
					pontoEmenda := len(resultado)
					esquerda := pontoEmenda - 1
					for esquerda >= 0 && ehVogal(novaJuncao[esquerda]) {
						esquerda--
					}
					esquerda++
					direita := pontoEmenda
					for direita < len(novaJuncao) && ehVogal(novaJuncao[direita]) {
						direita++
					}
					tamanhoBlocoVogais := direita - esquerda
					if tamanhoBlocoVogais >= 3 {
						resultado = resultado[:esquerda] + proxima
					} else {
						resultado = novaJuncao
					}
				}
			} else {
				resultado += " " + proxima
			}
		}
		fmt.Println(resultado)
	}
}