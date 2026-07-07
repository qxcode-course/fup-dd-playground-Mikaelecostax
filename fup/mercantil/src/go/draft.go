package main
import "fmt"
func main() {
    var rodadas int
    fmt.Scan(&rodadas)
    valorr := make([]float64, rodadas)
    chutesp := make([]float64, rodadas)
    escolhass := make([]string, rodadas)
    for i := 0; i < rodadas; i++ {
        fmt.Scan(&valorr[i])
    }
    for i := 0; i < rodadas; i++ {
        fmt.Scan(&chutesp[i])
    }
    for i := 0; i < rodadas; i++ {
        fmt.Scan(&escolhass[i])
    }
    vt := 0
    vs := 0
    for i := 0; i < rodadas; i++ {
        real:= valorr[i]
        chute:= chutesp[i]
        escolha:= escolhass[i]
    if chute == real {
        vt++
        continue
    }
    if real > chute {
        if escolha == "M"{
            vs++
        } else {
            vt++
        }
    }
    if real < chute {
        if escolha == "m"{
            vs++
        } else {
            vt++
        }
    }
}
     if vt > vs {
        fmt.Printf("primeiro\n")
    } else if vs > vt {
        fmt.Printf("segundo\n")
    } else {
        fmt.Printf("empate\n")
    }
}