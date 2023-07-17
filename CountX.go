package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the number to be multiplied: ")
	xStr := Userinput()
	x, _ := strconv.Atoi(xStr)
	fmt.Print("Enter the amount to multiply the number by: ")
	nStr := Userinput()
	n, _ := strconv.Atoi(nStr)
	result := make([]int, n)  // Make an array out of n(multiples)
	for i := 1; i <= n; i++ { // For n range
		result[i-1] = i * x // Use i to convey a multiplication, but also track the position.
	}
	fmt.Println(result)
}

func Userinput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}
