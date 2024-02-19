package main

import (
	"fmt"
)

func main() {

	var k1, m, k2, p2, n2 int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)
	//test case
	//M := 20 // Этажность
	//K2 := 41 // известн.: кв
	//P2 := 1  //           подъезд
	//N2 := 11 //           этаж
	//K1 := 89

	avgK := k2 / (n2 + (m * (p2 - 1))) // кол кв на этаже
	if avgK == 0 {
		fmt.Println("-1 -1")
	} else {
		Psearch := k1 / (m * avgK)

		if Psearch <= 0 || n2 <= 0 {
			fmt.Println(0, 0)
		}
		if m == 1 {
			fmt.Println(0, avgK)
		} else {
			fmt.Println(Psearch+1, avgK)
		}

	}
}
