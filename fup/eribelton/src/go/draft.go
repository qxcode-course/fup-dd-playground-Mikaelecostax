package main
import (
	"fmt"

)
func calculaAscologia(palavra string) int {
    soma := 0
    frequencia := make(map[byte]int)
    for i := 0; i < len(palavra); i++ {
        char := palavra[i]
        valorASCII := int(char)
        charMinusculo := char
        if char >= 'A' && char <= 'Z' {
            charMinusculo = char + 32
        }
        soma += (valorASCII - (50 * frequencia[charMinusculo]))
        frequencia[charMinusculo]++
    }
    return soma % 100
}
func ehVogalMaiuscula(c byte) bool {
    return c == 'A'
}
func main() {
	var palavraOriginal string
	if _, err := fmt.Scan(&palavraOriginal); err != nil {
		return
	}
	valorOriginal := calculaAscologia(palavraOriginal)
	menorValor := valorOriginal
	melhorLetra := ""
	for c := 'a'; c <= 'z'; c++ {
		palavraTestada := palavraOriginal + string(c)
		valorTestado := calculaAscologia(palavraTestada)
		if valorTestado < menorValor {
			menorValor = valorTestado
			melhorLetra = string(c)
		}
	}
	novaPalavra := palavraOriginal
	if melhorLetra != "" {
		novaPalavra = palavraOriginal + melhorLetra
	}
	fmt.Println(valorOriginal)
	fmt.Println(novaPalavra)
	fmt.Println(menorValor)
}