package main
import "fmt"
func main() {
    var nome string
    var idade int
    var classificacao string
    _, errnome := fmt.Scan(&nome)
    _, erridade := fmt.Scan(&idade)
    if errnome != nil || erridade != nil {
		return
	}
    if idade < 12 {
		classificacao = "crianca"
	} else if idade < 18 {
		classificacao = "jovem"
	} else if idade < 65 {
		classificacao = "adulto"
	} else if idade < 1000 {
		classificacao = "idoso"
	} else {
		classificacao = "mumia"
	}
    fmt.Printf("%s eh %s\n", nome, classificacao)
}
