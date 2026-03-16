package main
import "fmt"
func main() {
    var tempo, hora, minuto, segundo int
    fmt.Scan(&tempo)
    hora = tempo/3600
    minuto = (tempo%3600)/60
    segundo = (tempo%3600)%60
    fmt.Printf("%d:%d:%d\n", hora, minuto, segundo)

}
