package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    arena:= make([][]string, n)
    LL, CL := -1, -1
    for i := 0; i < n; i++ {
        arena[i] = make([]string, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&arena[i][j])
            if arena[i][j]== "L" {
                LL = i
                CL = j
            }
        }
    }
    PG := 0
    PC := 0
    for i := 0; i <n; i ++ {
        for j := 0; j < n ; j++ {
            if LL != -1 && (i == LL || j == CL){
                continue
            }
        
        if arena[i][j] == "G"{
            PG += 2
        } else if arena[i][j] == "C"{
            if i+j == n-1 {
                PC +=2
            } else {
                PC +=1
            }
        }
    } 
    }
    if PG > PC {
        fmt.Println("Gladiadores")
    } else if PC > PG {
        fmt.Println("Condenados a morte")
    } else {
        fmt.Println("Ninguem")
    }
}