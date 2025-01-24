package main

import (
	"fmt"
	"strconv"
)

func cpf_validator(cpf string) {
	num1 := [9]int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	sum1 := 0
	for i := 0; i < 9; i++ {
		integer, _ := strconv.Atoi(string(cpf[i]))
		sum1 += integer * num1[i]
	}

	var first_digit int
	rest1 := sum1 % 11
	if rest1 < 2 {
		first_digit = 0
	} else {
		first_digit = 11 - rest1
	}

	num2 := [10]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	sum2 := 0
	for i := 0; i < 10; i++ {
		integer, _ := strconv.Atoi(string(cpf[i]))
		sum2 += integer * num2[i]
	}

	var second_digit int
	rest2 := sum2 % 11
	fmt.Println(rest2)
	if rest2 < 2 {
		second_digit = 0
	} else {
		second_digit = 11 - rest2
	}

	first_digit_str := strconv.Itoa(first_digit)
	fmt.Println(first_digit_str)
	second_digit_str := strconv.Itoa(second_digit)
	fmt.Println(second_digit_str)

	if first_digit_str == string(cpf[9]) && second_digit_str == string(cpf[10]) {
		fmt.Printf("the cpf is valid!")
	} else {
		fmt.Printf("the cps is not valid.")
	}
}

func main() {
	cpf := "69079635073"
	cpf_validator(cpf)
}
