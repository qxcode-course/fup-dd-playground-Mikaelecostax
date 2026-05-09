package main
import "fmt"
func main() {
   var n int
	fmt.Scan(&n)
    fmt.Printf("[ ")
   
    for i := 0; i <= 10; i++ {
        if i == n {
			continue
		} 
        
        if i == 10 {
			fmt.Print("ceu ")
		} else {
			fmt.Printf("%d ", i)
		} 
    }
    fmt.Println("]")
}
