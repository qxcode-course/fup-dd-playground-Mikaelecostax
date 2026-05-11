package main
import "fmt"
func main() {
    var totalAlbum, totalPossui int
	if _, err := fmt.Scan(&totalAlbum); err != nil {
		return
	}
	if _, err := fmt.Scan(&totalPossui); err != nil {
		return
	}
	possui := make(map[int]int)
	for i := 0; i < totalPossui; i++ {
		var fig int
		fmt.Scan(&fig)
		possui[fig]++
	}
	fmt.Print("[")
	for i := 1; i <= 100; i++ {
		for j := 0; j < possui[i]-1; j++ {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println(" ]")
	fmt.Print("[")
	for i := 1; i <= totalAlbum; i++ {
		if possui[i] == 0 {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println(" ]")
}
