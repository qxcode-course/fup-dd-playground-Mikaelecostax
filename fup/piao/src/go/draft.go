package main
import "fmt"
func main() {
   var limite, n int
	fmt.Scan(&limite, &n)
	ganhador := -1
	perdedor := 0
	menorDistanciaValida := 999999
	maiorDistanciaGeral := -1
	for i := 0; i < n; i++ {
		var jogada int
		fmt.Scan(&jogada)
		distanciaAbsoluta := jogada
		if distanciaAbsoluta < 0 {
			distanciaAbsoluta = -distanciaAbsoluta
		}
		if distanciaAbsoluta <= limite {
			if distanciaAbsoluta <= menorDistanciaValida {
				menorDistanciaValida = distanciaAbsoluta
				ganhador = i
			}
		}
		if distanciaAbsoluta >= maiorDistanciaGeral {
			maiorDistanciaGeral = distanciaAbsoluta
			perdedor = i
		}
        }
	if ganhador == -1 {
		fmt.Println("nenhum")
	} else {
		fmt.Println(ganhador)
	}
	fmt.Println(perdedor)
}
