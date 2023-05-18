package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			total := (int(num1[i]-'0') * int(num2[j]-'0'))
			result[i+j+1] += total
			result[i+j] += result[i+j+1] / 10
			result[i+j+1] = result[i+j+1] % 10
		}
	}

	var res string
	flag := false

	for i := 0; i < len(num1)+len(num2); i++ {
		if result[i] == 0 && flag == false {
			continue
		}
		if result[i] != 0 || flag == true {
			flag = true
			res += strconv.Itoa(result[i])
		}
	}
	return res
}

func main() {
	a := "999"
	b := "0"
	fmt.Println(multiply(a, b))
}
