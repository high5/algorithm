
package main

import (
    "fmt"
    "math"
    //"strconv"
    //"reflect"
)

/*
* F(0) = 0
* F(1) = 1
* F(n) = F(n-1) + F(n-2) 
*/
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1)+fib1(n-2)
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2_sub(1, 1, n)
}

func fib2_sub(a, b, c int) int {
	if c <= 2 {
		return a
	}
	return fib2_sub(a+b, a, c-1) 
}

func fib3(n int) int{
	fib0 := 0;
    fib1 := 1;
    ret := 0;
    for i :=0; i<n; i++{
        ret = fib1
        fib1 = fib0 + fib1;
        fib0 = ret;
    }
    return ret;
}

func fib4(n float64) int {
	return int(math.Floor(math.Pow((1+math.Sqrt(5))/2, n) / math.Sqrt(5) + 1/2))
}


func main() {
	fmt.Println(fib1(38))
	fmt.Println(fib2(38))
	fmt.Println(fib3(38))
	fmt.Println(fib4(38))
}












