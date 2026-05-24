package main
import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		resultado := make([]rune, 0, len(texto))
		for _, caractere := range texto {
			if unicode.IsUpper(caractere) {
				resultado = append(resultado, unicode.ToLower(caractere))
			} else if unicode.IsLower(caractere) {
				resultado = append(resultado, unicode.ToUpper(caractere))
			} else {
				resultado = append(resultado, caractere)
			}
		}
		fmt.Println(string(resultado))
	}
}
