//Определите является ли билет счастливым. Счастливым, считается билет, в шестизначном номере которого сумма первых трех цифр совпадает с суммой трех последних формат входных данных на вход подается номер билета - одно шестизначное число. формат выходных данных. Выведите "yes", если билет счастливый, в противном случае - "no"
package main

import "fmt"

func main() {
	CheckTicket()
}


func CheckTicket() {
	ticket := 409535

	d1 := ticket / 100000       
	d2 := (ticket / 10000) % 10 
	d3 := (ticket / 1000) % 10  

	d4 := (ticket / 100) % 10  
	d5 := (ticket / 10) % 10    
	d6 := ticket % 10          

	firstSum := d1 + d2 + d3
	lastSum := d4 + d5 + d6

	if firstSum == lastSum {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
