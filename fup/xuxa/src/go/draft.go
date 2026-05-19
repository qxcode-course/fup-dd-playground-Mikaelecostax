package main
import ("fmt"
    "bufio"
	"os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		frase := scanner.Text()
		caracteres := []rune(frase)
		for i := len(caracteres) - 1; i >= 0; i-- {
			fmt.Print(string(caracteres[i]))
		}
		fmt.Println()
	}
}
