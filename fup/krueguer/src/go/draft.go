package main

import (
	"bufio"
	"fmt"
	"strings"
    "os"
)
func main() {
    var n int
    fmt.Scan(&n)
    scanner := bufio.NewScanner(os.Stdin)
    for i := 0; i < n; i++ {
        scanner.Scan()
        frase := scanner.Text()
        MS := ""
        SA := ""
        for _, letra := range frase {
            if strings.ContainsRune("aeiou", letra){
                SA += string(letra)
            } else {
                if len(SA) > len(MS) {
                    MS = SA
                }
                SA = ""
            }
        }
        if len(SA) > len(MS){
            MS = SA
        }
        fmt.Println(MS)
    }
}