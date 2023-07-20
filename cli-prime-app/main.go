package main

import "fmt"

func main() {
	_, res := IsPrime(5)
	fmt.Print(res)
}

func IsPrime(num int) (bool, string) {
	if num == 0 || num == 1 || num <= 0 {
		return false, fmt.Sprintf("%v is not a prime number\n", num)
	}

	var i int
	for i = 2; i <= num/2; i++ {
		if num%i == 0 {
			return false, fmt.Sprintf("%v is not a prime number, its divisable by %v \n", num, i)
		}
	}
	return true, fmt.Sprintf("%v is a prime number \n", num)
}
