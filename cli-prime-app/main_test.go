package main

import (
	"fmt"
	"testing"
)

func Test_IsPrime(t *testing.T) {
	num := 0
	res, msg := IsPrime(num)
	if res != false {
		t.Errorf("With %d as test param, result is prime, but its expected to be not prime", num)
	}

	if msg != fmt.Sprintf("%v is not a prime number\n", num) {
		t.Errorf("With %d as test param, response msg is not the same as the expected one !", num)
	}
}
