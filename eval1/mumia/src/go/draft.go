package main
import "fmt"
func main() {
    var mario string
    var x int
    var jose string
    var y int
    fmt.Scan(&x, &mario, &jose, &y)
    if x < 12  {
   fmt.Println("mario eh crianca")
} else if x < 18 {
   fmt.Println("mario eh jovem")
} else if x < 65 {
   fmt.Println("mario eh adulto")
} else if x < 1000 {
  fmt.Println("jose eh idoso")
} 
} 