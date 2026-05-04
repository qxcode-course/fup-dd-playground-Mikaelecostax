package main
import "fmt"
func main() {
    var salario float64
    var novosalario float64
    fmt.Scan(&salario)
    if salario <= 1000.00 {
		novosalario = salario * 1.20
	} else if salario <= 1500.00 {
		novosalario = salario * 1.15
	} else if salario <= 2000.00 {
		novosalario = salario * 1.10
	} else {
		novosalario = salario * 1.05
	}
    fmt.Printf("%.2f\n", novosalario)
}
