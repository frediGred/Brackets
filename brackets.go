package main

import (
	"strings"
)

func getKeyMap(symb string, mapping map[string]string) string {
	for key, _ := range mapping {
		if symb == key {
			return key
		}
	}
	return ""
}

func CorrectParentheses(n []string) []bool { //функция принимает массив строк и возвращает массив boolean
	Bracket := map[string]string{"}": "{", "]": "[", ")": "("}
	var mp []bool

	for i, _ := range n { // проход по масиву
		str := strings.Split(n[i], "") //разбиваем строку на массив чтобы можно было пройти по каждому элементу
		iterate := 1                   // переменная для обратного поиска
		var array []string             // массив
		for j, symbol := range str {
			if symbol == "(" || symbol == "{" || symbol == "[" {
				array = append(array, symbol)
				continue
			} else if j-iterate < 0 {
				mp = append(mp, false)
				break
			} else if symbol == getKeyMap(symbol, Bracket) {
				if str[j-iterate] == Bracket[symbol] {
					array = append(array, symbol)
					iterate += 2
					continue
				}
			} else {
				mp = append(mp, false)
				break
			}
		}
		if strings.Join(str, "") == strings.Join(array, "") {
			mp = append(mp, true)

		} else {
			mp = append(mp, false)
		}
	}
	return mp

}
