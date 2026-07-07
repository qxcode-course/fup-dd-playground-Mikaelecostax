package main
import( "fmt"
    "bufio"
    "os"
     "unicode")
func main() {
   var n int
   fmt.Scanln(&n)
   scanner := bufio.NewScanner(os.Stdin)
   for i := 0; i < n; i++ {
    scanner.Scan()
    frase := []rune(scanner.Text())
    maiuscula := false
    for _, caractere := range frase{
        if caractere != ' '{
            maiuscula = unicode.IsUpper(caractere)
            break
        }
    }
    for i, caractere := range frase {
        if caractere != ' '{
            if maiuscula {
                frase[i] = unicode.ToUpper(caractere)
            } else {
                frase[i] = unicode.ToLower(caractere)
            }
            maiuscula = !maiuscula
        }
    }
        fmt.Println(string(frase))
   }
}