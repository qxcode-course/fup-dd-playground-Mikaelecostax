package main
import ("fmt"
    "strings" )
func main() {
    var x int
    fmt.Scan(&x)
    for i := 0; i < x; i++ {
        var codigoUl, codigop string
        fmt.Scan(&codigoUl, &codigop)
        ultronL := strings.ToLower(codigoUl)
        pessoaL := strings.ToLower(codigop)
        letrasUl := make(map[rune]bool)
        for _, letra := range ultronL {
            letrasUl[letra]= true
        }
        letrascorrespondentes := 0
        for _, letra := range pessoaL {
            if letrasUl[letra] {
                letrascorrespondentes++
            }
        }
        totalletrasp := len(pessoaL)
        if letrascorrespondentes == totalletrasp{
            fmt.Println("chefe")
        }else if letrascorrespondentes*2 > totalletrasp{
            fmt.Println("ultron")
        } else {
            fmt.Println("pessoa")
        }
    }
}