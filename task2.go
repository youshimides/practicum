package main

import "fmt"

func main() {

	Task2()
}

// по данному трехзначному числу определите все ли его цифры различны. Формат входных данных на вход подается одно натуральное трехзначное число. Формат выходных данных Выведите "YES", если все цифры различны, в противном - "NO"
func Task2() {
	num := 123

	hundreds := num / 100   
	tens := (num / 10) % 10 
	ones := num % 10        

	if hundreds != tens && hundreds != ones && tens != ones {
		fmt.Println("YES")
	} else {
		fmt.Println("No")
	}
}