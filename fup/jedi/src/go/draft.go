package main
import "fmt"
func main() {
    var numet int
    fmt.Scan(&numet)
    Jedi := 0
    Sith := 0
    metade := numet/2 
    for i := 0; i < numet; i++ {
        var forca int 
        fmt.Scan(&forca)
     if i < metade {
        Jedi += forca
    } else { 
        Sith += forca
    }
    }
    if Jedi > Sith {
        fmt.Println("Jedi")
    } else if Sith > Jedi {
        fmt.Println("Sith")
    } else {
        fmt.Println("Empate")
    }
}