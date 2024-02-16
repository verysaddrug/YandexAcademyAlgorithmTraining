package main

import (
	"fmt"
)

func main() {
	// fmt.Println(10 / 3)
	var K1, M, K2, P2, N2 int        // Range: -2147483648 through 2147483647.
	fmt.Scan(&K1, &M, &K2, &P2, &N2) // Квартира, кол-во этажей, кв, подъезд, этаж
	//	  89  20   41    1   11
	fmt.Println(restoreInfo(K1, M, K2, P2, N2))
}

func restoreInfo(K1, M, K2, P2, N2 int) (P1, N1 int) {
	if N2 > M {
		return
	}
	P1, N1 = -1, -1
	var y int

	for i := 1; i <= 1000; i++ {
		y = K2
		y -= i * M * (P2 - 1)
		y -= i * (N2 - 1)
		if y < 1 || y > i {
			continue
		}

		// x = K1
		a := K1/(i*M) + 1
		if K1%(i*M) == 0 {
			a--
		}
		b := K1 - (a-1)*i*M
		b /= i
		b++

		if (K1-((a-1)*i*M))%i == 0 {
			b--
		}
		if a < 1 || b < 1 {
			continue
		}
		c := K1 - (a-1)*i*M - (b-1)*i
		// cout << i << ' ' << b <<endl;
		if c < 1 || c > i {
			continue
		}

		if P1 == -1 {
			P1 = a
		} else if P1 != a {
			P1 = 0
		}

		if N1 == -1 {
			N1 = b
		} else if N1 != b {
			N1 = 0
		}
	}
	return
}
