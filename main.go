package main

import "fmt"



func main() {
	var brackets []string = []string{"({[]})", "(({}))", "()", "()))", "(){}[]"}
	rez := CorrectParentheses(brackets)
	for i, value := range rez {
		fmt.Printf("Brackets:%d - %t \n", i, value)
	}
}