package main

import "fmt"
import "time"

func main() {

	// ======= SIMPLE SWITCH ========
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("Satu")
	// case 2:
	// 	fmt.Println("Dua")
	// case 3:
	// 	fmt.Println("Tiga")
	// default:
	// 	fmt.Println("Bukan 1, 2, atau 3")
	// }


	// ======= SWITCH WITH MULTIPLE CASES ========

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	case time.Friday:
		fmt.Println("Sholat jumat")
	default:
		fmt.Println("Kerja kocak")
	}
}
