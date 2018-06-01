package main

import "fmt"

func main() {
	fmt.Println("hello")
	twoDigit()
	threeDigit()
}

func twoDigit() {
	fmt.Println("two digits wide")
	result := float64(1) / float64(9899)
	// only accurate up to 9 segments
	// need to use math lib for accuracy
	printItOut(result, 9, 2)
}

func threeDigit() {
	fmt.Println("three digits wide")
	result := float64(1) / float64(998999)
	// only accurate up to 7 segments
	// need to use math lib for accuracy
	printItOut(result, 12, 3)
}

func printItOut(in float64, n, wide int) {
	for i := 0; i < n; i++ {
		set := fmt.Sprintf("%."+fmt.Sprint((i+1)*wide)+"f", in)
		fmt.Println(set[len(set)-wide:])
	}
}
