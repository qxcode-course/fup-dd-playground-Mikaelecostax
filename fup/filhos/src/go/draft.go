package main
import "fmt"
func main() {
    var idademaisnovo int
	var quantidadefilhos int
    _, _ = fmt.Scan(&idademaisnovo, &quantidadefilhos)
    idadeatual := idademaisnovo
    for i := 0; i < quantidadefilhos; i++ {
		fmt.Println(idadeatual)
        idadeatual += 2
    }
}
