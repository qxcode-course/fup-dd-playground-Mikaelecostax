package main
import( "fmt"
    "io")
func main() {
    var capacidade, movimentacao int
    fmt.Scan(&capacidade)
	passageirosatuais := 0
    for {
        _, err := fmt.Scan(&movimentacao)
		if err == io.EOF {
			break
        } 
        passageirosatuais += movimentacao
		if passageirosatuais >= 2*capacidade {
			fmt.Println("hora de partir")
			return
		}
		if passageirosatuais <= 0 {
			fmt.Println("vazio")
			passageirosatuais = 0
		} else if passageirosatuais < capacidade {
			fmt.Println("ainda cabe")
		} else {
			fmt.Println("lotado")
        }
    }
}