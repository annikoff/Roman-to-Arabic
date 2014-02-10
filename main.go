package main

import (
	"fmt"
)

func main() {
	fmt.Println("Roman Arabic Numerals Converter")
	var input, rn string
	var an int
	fmt.Println("1. R to A")
	fmt.Println("2. A to R")
	fmt.Println("0. Exit")
	for {
		fmt.Print("Enter command: ")
		fmt.Scan(&input)
		fmt.Println("---------")
		switch input {
		case "1":
			fmt.Println("0. Back to menu")
			for {
				fmt.Print("Enter a Roman Number: ")
				fmt.Scan(&rn)
				if rn == "0" {
					break
				}
				fmt.Println(rn, "is", r2a(rn))
				fmt.Println("---------")
			}
			break
		case "2":
			fmt.Println("0. Back to menu")
			for {
				fmt.Print("Enter an Arabic Number: ")
				fmt.Scan(&an)
				if an == 0 {
					break
				}
				fmt.Println(an, "is", a2r(an))
				fmt.Println("---------")
			}
			break
		case "0":
			return
		}
	}
}

//Roman Numerals to Arabic Numerals
func r2a(romanNum string) int {
	num := 0
	currNum := 0
	lastNum := 0
	for i := 0; i < len(romanNum); i++ {
		switch string(romanNum[i]) {
		case "I":
			currNum = 1
			break
		case "V":
			currNum = 5
			break
		case "X":
			currNum = 10
			break
		case "L":
			currNum = 50
			break
		case "C":
			currNum = 100
			break
		case "D":
			currNum = 500
			break
		case "M":
			currNum = 1000
			break
		}
		if lastNum != 0 && currNum > lastNum {
			num += currNum - (lastNum * 2)
		} else {
			num += currNum
		}
		lastNum = currNum
	}
	return num
}

//Arabic Numerals to Roman Numerals
func a2r(arabicNum int) string {
	roman := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabic := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for i := 0; i < len(roman); i++ {
		for arabicNum >= arabic[i] {
			arabicNum -= arabic[i]
			result += roman[i]
		}
	}
	return result
}
