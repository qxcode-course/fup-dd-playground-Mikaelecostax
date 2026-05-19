package main
import ( "fmt"
    "unicode"
)
func main() {
   var c rune
	fmt.Scanf("%c", &c)
	if unicode.IsUpper(c) {
		c = unicode.ToLower(c)
	} else if unicode.IsLower(c) {
		c = unicode.ToUpper(c)
	}
	fmt.Printf("%c\n", c)
}
