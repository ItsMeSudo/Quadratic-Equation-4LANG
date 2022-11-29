package main

import (
	"fmt"
	"math"
)

func main() {
	var a_val float64
	var b_val float64
	var c_val float64
	fmt.Printf("Add meg az A értékét: ")
	fmt.Scanln(&a_val)
	fmt.Printf("")
	fmt.Printf("Add meg az B értékét: ")
	fmt.Scanln(&b_val)
	fmt.Printf("")
	fmt.Printf("Add meg az C értékét: ")
	fmt.Scanln(&c_val)
	fmt.Printf("")
	fmt.Printf("A értéke: %v, B értéke: %v, C értéke: %v\n", a_val, b_val, c_val)
	var diszkriminans float64 = (b_val * b_val) - (4 * a_val * c_val)
	fmt.Printf("A diszkrimináns értéke: %v\n", diszkriminans)
	if diszkriminans < 0 {
		fmt.Printf("Az egyenletnek nincs megoldása a valós számok halmazán!\n")
		return
	} else if diszkriminans == 0 {
		fmt.Printf("Az egyenletnek két egyenlő (kettős gyöke) van!\n")
	} else if diszkriminans > 0 {
		fmt.Printf("Az egyenletnek két különböző valós gyöke van\n")
	} else {
		fmt.Printf("HIBA\n")
		return
	}
	var res_1 float64 = -b_val + math.Sqrt(float64(diszkriminans))
	var res_2 float64 = -b_val - math.Sqrt(float64(diszkriminans))
	var result_1 float64 = (res_1 / (2 * a_val))
	var result_2 float64 = (res_2 / (2 * a_val))
	fmt.Printf("X1 értéke: %v\n", result_1)
	fmt.Printf("X2 értéke: %v\n", result_2)

	fmt.Printf("SUDO#0814 2022")

}
