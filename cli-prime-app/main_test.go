package main

import (
	"fmt"
	"testing"
)

type test_cases struct {
	title           string
	val             int
	expected_msg    string
	expected_result bool
}

func Test_IsPrime(t *testing.T) {
	tests := []test_cases{
		{title: "testing 7", val: 7, expected_msg: fmt.Sprintf("%v is a prime number \n", 7), expected_result: true},
		{title: "testing 0", val: 0, expected_msg: fmt.Sprintf("%v is not a prime number\n", 0), expected_result: false},
		{title: "testing 1", val: 1, expected_msg: fmt.Sprintf("%v is not a prime number\n", 1), expected_result: false},
		{title: "testing 8", val: 8, expected_msg: fmt.Sprintf("%v is not a prime number, its divisable by %v \n", 8, 2), expected_result: false},
	}
	for _, entry := range tests {
		result, msg := IsPrime(entry.val)
		if msg != entry.expected_msg {
			t.Errorf("%s : With %d as test param, response msg is not the same as the expected one !", entry.title, entry.val)
		}
		if (result == true && entry.expected_result == false) || (result == false && entry.expected_result == true) {
			t.Errorf("%s : With %d as test param, result is prime, but its expected to be not prime", entry.title, entry.val)

		}
	}
}
