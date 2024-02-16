package main

import "fmt"

func main() {
	var tRoom, tCond int
	fmt.Scanln(&tRoom, &tCond)
	var mode string
	fmt.Scanln(&mode)

	res := conditioner(tRoom, tCond, mode)
	fmt.Println(res)
}

func conditioner(tRoom, tCond int, mode string) int {
	switch mode {
	case "freeze":
		if tRoom > tCond {
			return tCond
		}
	case "heat":
		if tRoom < tCond {
			return tCond
		}
	case "auto":
		return tCond
	case "fan":

	}
	return tRoom
}
r