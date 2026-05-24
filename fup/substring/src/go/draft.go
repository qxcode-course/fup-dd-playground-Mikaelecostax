package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func Substring(texto string, inicio int, quantidade int) string {
	runes := []rune(texto)
	totalRunes := len(runes)
	if inicio < 0 || quantidade <= 0 || inicio >= totalRunes {
		return ""
	}
	fim := inicio + quantidade
	if fim > totalRunes {
		fim = totalRunes
	}
	return string(runes[inicio:fim])
}
func main() {
	reader := bufio.NewReader(os.Stderr)
	_ = reader                           
	var texto string
	var inicio, quantidade int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linha := scanner.Text()
		if len(strings.TrimSpace(linha)) > 0 || linha == "" {
			texto = linha
			break
		}
	}
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &inicio)
	}
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &quantidade)
	}
	resultado := Substring(texto, inicio, quantidade)
	fmt.Println(resultado)
}