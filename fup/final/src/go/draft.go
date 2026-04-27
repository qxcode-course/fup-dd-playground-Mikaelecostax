package main
import "fmt"
func main() {
    var nota1, nota2, notafinal, media, mediafinal int
    fmt.Scan(&nota1, &nota2, &notafinal)
    media = (nota1 + nota2) / 2
    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else {
        mediafinal = (media + notafinal) / 2
    
    if mediafinal >=5 {
        fmt.Println("aprovado na final")
    } else  {
        fmt.Println("reprovado na final")
    }
    }
}
