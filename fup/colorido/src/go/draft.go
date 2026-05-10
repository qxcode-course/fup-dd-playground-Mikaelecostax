package main
import ("fmt"
    "strings")
func main() {
    var npedra int
	var peinicial string
	fmt.Scan(&npedra, &peinicial)
    var resultado []string
	peatual := peinicial
	for i := 0; i <= 10; i++ {
		if i == npedra {
			continue
		}
        if i == 10 {
			resultado = append(resultado, "ceu")
		} else {
			resultado = append(resultado, fmt.Sprintf("%d%s", i, peatual))
		}
        if peatual == "d" {
			peatual = "e"
		} else {
			peatual = "d"
		}
    }
    fmt.Printf("[ %s ]\n", strings.Join(resultado, " "))
}
