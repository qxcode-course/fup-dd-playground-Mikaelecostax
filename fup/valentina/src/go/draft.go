package main
import (
	"fmt"
)
func main() {
	var c1Char, operacaoChar, c2Char string
	if _, err := fmt.Scan(&c1Char); err != nil {
		return
	}
	if _, err := fmt.Scan(&operacaoChar); err != nil {
		return
	}
	if _, err := fmt.Scan(&c2Char); err != nil {
		return
	}
	c1 := c1Char[0]
	op := operacaoChar[0]
	c2 := c2Char[0]
	val1 := int(c1 - 'a')
	val2 := int(c2 - 'a')
	var resultadoNum int
	if op == '+' {
		resultadoNum = (val1 + val2) % 26
	} else if op == '-' {
		resultadoNum = (val1 - val2 + 26) % 26
	}
	resultadoLetra := byte(resultadoNum) + 'a'
	fmt.Printf("%c\n", resultadoLetra)
}
