package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Input the first number: ")
	num1 := getNum()
	fmt.Println("Input the second number: ")
	num2 := getNum()
	sum := num1 + num2
	fmt.Println("The sum is " + strconv.Itoa(sum))
}

func getNum() int {
	text := retrieve()
	num, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Please input a valid number!")
		return getNum()
	}

	return num
}

func retrieve() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	err := scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	return scanner.Text()
}
