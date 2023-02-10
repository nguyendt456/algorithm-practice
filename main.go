package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	var result string
	charJump := numRows*2 - 2
	if charJump == 0 {
		return s
	}
	for row := 0; row < numRows; row++ {
		for nextZigZacChar := 0; nextZigZacChar+row < len(s); {
			result += string(s[nextZigZacChar+row])
			if row == 0 || row == numRows-1 {
				nextZigZacChar += charJump
				continue
			}
			if nextZigZacChar+row+(numRows-row)*2-2 >= len(s) {
				break
			}
			result += string(s[nextZigZacChar+row+(numRows-row)*2-2])
			nextZigZacChar += charJump
		}
	}
	return result
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))

}
