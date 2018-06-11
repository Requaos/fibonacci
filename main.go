package main

import (
	"fmt"
	"math/big"
)

type iterativeState struct {
	first, second int
}

func main() {
	fmt.Println("Fibonacci patterns:")

	// Fraction derived and shaped with format library
	twoDigit()
	threeDigit()

	// Impractical recursive calculation
	for i := 0; i < 5; i++ {
		fmt.Println(recursiveFibonacci(i))
	}

	// Slightly better iterative calculation
	iterativeFibonacci(10)

}

func twoDigit() {
	fmt.Println("two digits wide")
	result := new(big.Float).Quo(big.NewFloat(1), big.NewFloat(9899))
	// only accurate up to 9 segments
	// need to use math lib for accuracy
	printItOut(*result, 12, 2)
}

func threeDigit() {
	fmt.Println("three digits wide")
	result := new(big.Float).Quo(big.NewFloat(1), big.NewFloat(998999))
	// only accurate up to 7 segments
	// need to use math lib for accuracy
	printItOut(*result, 12, 3)
}

func printItOut(in big.Float, n, wide int) {
	floater, _ := in.Float64()
	for i := 0; i < n; i++ {
		set := fmt.Sprintf("%."+fmt.Sprint((i+1)*wide)+"f", floater)
		fmt.Println(set[len(set)-wide:])
	}
}

func recursiveFibonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}

func iterativeFibonacci(n int) {
	fmt.Println(0)
	v := iterativeState{0,1}
	for i := 1; i < n; i++ {
		fmt.Println(v.second)
		v = iterativeState{v.second, v.first + v.second}
	}
}