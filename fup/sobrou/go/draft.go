package main
import "fmt"
func main() {
    var aq, bq, cq float64
    var dv, ev, fv float64
    fmt.Scan(&aq, &bq, &cq)
    fmt.Scan(&dv, &ev, &fv)
    total := aq*dv + bq*ev + cq*fv
    fmt.Printf("%.2f\n", total)
    
}
