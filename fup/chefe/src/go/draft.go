package main
import ("fmt"
    "strings" )
func main() {
    var cou string
    var pessoa string
    fmt.Scanln(&cou)
    ultronl := strings.ToLower(cou)
    letrasul := make(map[rune]bool)
    for _, letra := range ultronl{
        letrasul[letra] = true
    }
    primeiro := true
    for {
        _, err := fmt.Scan(&pessoa)
        if err != nil {
            break
        }
        pL := strings.ToLower(pessoa)
        letrascorrespondentes := 0 
        for _, letra := range pL {
            if letrasul[letra]{
                letrascorrespondentes ++
            }
        }
        totalLP := len(pessoa)
        var resultado string
        if letrascorrespondentes == totalLP {
            resultado = "chefe"
        } else if letrascorrespondentes*2 > totalLP {
            resultado = "ultron"
        } else {
            resultado = "pessoa"
        }
        if primeiro {
            fmt.Print(resultado)
            primeiro = false 
        } else {
            fmt.Print(" " + resultado)
        }
    } 
        fmt.Println()
}