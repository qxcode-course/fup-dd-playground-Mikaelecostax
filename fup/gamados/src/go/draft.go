package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		palavras := strings.Fields(texto)
		if len(palavras) <= 1 {
			fmt.Println("sim")
			return
        }
		ordenado := true
		for i := 0; i < len(palavras)-1; i++ {
			if palavras[i] > palavras[i+1] {
				ordenado = false
				break
			}
		}
		if ordenado {
			fmt.Println("sim")
		} else {
			fmt.Println("nao")
		}
	}
}
