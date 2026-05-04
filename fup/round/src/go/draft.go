package main
import ( "fmt"
        "math"
)
func main() {
    var op string
    var num float64
    _, err := fmt.Scan( &op, &num)
	if err != nil {
		return
	}
    var resultado int
    switch op {
	case "c":
		resultado = int (math.Ceil(num))
	case "f":
		resultado = int (math.Floor(num))
	case "r":
		resultado = int (math.Round(num))
	}
    fmt.Println(resultado)
}
