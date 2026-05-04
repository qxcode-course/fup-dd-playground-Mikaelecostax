package main
import "fmt"
func main() {
    var jog1 string
    var jog2 string
    fmt.Scan(&jog1, &jog2)
	if jog1 == jog2 {
		fmt.Println("empate")
        return
	} 
    if (jog1 == "R" && jog2 == "S") ||
		(jog1 == "S" && jog2 == "P") ||
		(jog1 == "P" && jog2 == "R") {
		fmt.Println("jog1")
	}else  {
		fmt.Println("jog2")
	}
}
