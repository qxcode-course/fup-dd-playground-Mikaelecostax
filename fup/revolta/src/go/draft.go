package main
import "fmt"
func main() {
   var n int
	fmt.Scan(&n)
	somaSoldados := 0 
	somaRebeldes := 0 
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x%2 != 0 {
			somaSoldados += x
		} else {
			somaRebeldes += x
		}
	}
	if somaSoldados > somaRebeldes {
		fmt.Println("soldados")
	} else if somaRebeldes > somaSoldados {
		fmt.Println("rebeldes")
	} else {
		fmt.Println("empate")
	}
}
