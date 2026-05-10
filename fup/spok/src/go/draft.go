package main
import "fmt"
func main (){
    var id int
	fmt.Scan(&id)
	original := id
	invertido := 0
	for temp := id; temp > 0; temp /= 10 {
		ultimodigito := temp % 10
		invertido = (invertido * 10) + ultimodigito
	}
	if original == invertido {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}