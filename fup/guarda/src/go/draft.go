package main
import "fmt"
func main() {
    var wifi int
    var login int
    var admin int
    fmt.Scan(&wifi, &login, &admin)
    if wifi == 0 { 
        fmt.Println("you must connect to wifi")
    } else if login == 0 {
         fmt.Println("you need to login first")
    } else if admin == 0 {
         fmt.Println("you must to login as admin")
    } else if wifi == 1 && login == 1 && admin == 1 {
         fmt.Println("done")
    }
   
}
