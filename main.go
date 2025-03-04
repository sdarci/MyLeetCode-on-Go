package main

import (
	"fmt"
	"leetcode/stringToInteger"
	"time"
)

func main() {

	startTime := time.Now()
	// atoi
	println(stringToInteger.MyAtoi("42"))
	elapsed := time.Since(startTime)
	fmt.Printf("MyAtoi function took %s \n", elapsed)
	startTime = time.Now()
	println(stringToInteger.MyAtoi("-42"))
	elapsed = time.Since(startTime)
	fmt.Printf("MyAtoi function took %s \n", elapsed)
	startTime = time.Now()
	println(stringToInteger.MyAtoi("       -42"))
	elapsed = time.Since(startTime)
	fmt.Printf("MyAtoi function took %s \n", elapsed)
	startTime = time.Now()
	println(stringToInteger.MyAtoi(" 0-1"))
	elapsed = time.Since(startTime)
	fmt.Printf("MyAtoi function took %s \n", elapsed)
	startTime = time.Now()
}
