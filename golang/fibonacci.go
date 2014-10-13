
package main

import (
    "fmt"
)

/*
* F(0) = 0
* F(1) = 1
* F(n) = F(n-1) + F(n-2) 
*/


func main() {
	fmt.Println(fib1(38))
}

func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1)+fib1(n-2)
}