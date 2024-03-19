package main

import "math/big"

func addBinary(a string, b string) string {
	x, _ := new(big.Int).SetString(a, 2)
	y, _ := new(big.Int).SetString(b, 2)

	return new(big.Int).Add(x, y).Text(2)
}
