package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func ehVogal(r rune) bool {
	c := strings.ToLower(string(r))
	return c == "a" || c == "e" || c == "i" || c == "o" || c == "u"
}
func ehConsoante(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return !ehVogal(r)
	}
	return false
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		var vogais []rune
		var consoantes []rune
		for _, caractere := range texto {
			if ehVogal(caractere) {
				vogais = append(vogais, caractere)
			} else if ehConsoante(caractere) {
				consoantes = append(consoantes, caractere)
			}
		}
		fmt.Println(string(vogais))
		fmt.Println(string(consoantes))
	}
}