package main
import "fmt"
func main() {
    var aq, bq, cq float64
    fmt.Scan(&aq, &bq, &cq)
    var dv, ev, fv float64
    fmt.Scan(&dv, &ev, &fv)
    var mk float64
    fmt.Scan(&mk)
    var tudo float64
    tudo = (aq*dv) + (bq*ev) + (cq*fv)
    troco := mk - tudo
    fmt.Printf("%.2f\n", troco)
    
}
