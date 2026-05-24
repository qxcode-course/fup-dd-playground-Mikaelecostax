package main
import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)
func ehVogal(r rune) bool {
	rMinuscula := unicode.ToLower(r)
	return rMinuscula == 'a' || rMinuscula == 'e' || rMinuscula == 'i' || rMinuscula == 'o' || rMinuscula == 'u'
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		resultado := make([]rune, 0, len(texto))
		for _, caractere := range texto {
			if unicode.IsLetter(caractere) {
				if ehVogal(caractere) {
					resultado = append(resultado, 'v')
				} else {
					resultado = append(resultado, 'c')
				}
			} else {
				resultado = append(resultado, caractere)
			}
		}
		fmt.Println(string(resultado))
	}
}