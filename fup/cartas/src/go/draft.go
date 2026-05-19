package main
import ("fmt"
    "strconv"
)
func main() {
   var n int
	fmt.Scan(&n)
	mao := make([]string, n)
	for i := 0; i < n; i++ {
		var carta int
		fmt.Scan(&carta)
		switch carta {
		case 1:
			mao[i] = "A"
		case 11:
			mao[i] = "J"
		case 12:
			mao[i] = "Q"
		case 13:
			mao[i] = "K"
		default:
			mao[i] = strconv.Itoa(carta)
		}
	}
	fmt.Print("[")
	for i, carta := range mao {
		fmt.Print(carta)
		if i < n-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
