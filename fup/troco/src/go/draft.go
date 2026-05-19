package main
import ( "fmt"
    "math"
)
type ItemTroco struct {
	valorCentavos int
	descricao     string
}
func main() {
    var valorOriginal float64
	fmt.Scan(&valorOriginal)
	trocoCentavos := int(math.Round(valorOriginal * 100))
	opcoes := []ItemTroco{
		{10000, "100.00"},
		{5000, "50.00"},
		{2000, "20.00"},
		{1000, "10.00"},
		{500, "5.00"},
		{200, "2.00"},
		{100, "1.00"},
		{50, "0.50"},
		{25, "0.25"},
		{10, "0.10"},
		{5, "0.05"},
	}
	for _, item := range opcoes {
		quantidade := trocoCentavos / item.valorCentavos
		for i := 0; i < quantidade; i++ {
			fmt.Printf("%d de %s\n", quantidade, item.descricao)
			break 
		}
		trocoCentavos %= item.valorCentavos
	}
	for i := 0; i < trocoCentavos; i++ {
		fmt.Printf("Falta %.2f\n", float64(trocoCentavos)/100.0)
		break
	}
}
