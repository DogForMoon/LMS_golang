package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func do_exp(num1, num2 float64, oper string) float64 {
	switch {
	case oper == "*":
		return num1 * num2
	case oper == "/":
		return num1 / num2
	case oper == "+":
		return num1 + num2
	case oper == "-":
		return num1 - num2
	default:
		return 0.0
	}
}

func do_checks(expression, r1 string, i int) (float64, error) {
	exp1, err1 := strconv.ParseFloat(expression[:i], 64)
	if err1 != nil {
		exp1, err1 = Calc(expression[:i])
	}
	exp2, err2 := strconv.ParseFloat(expression[i+1:], 64)
	if err2 != nil {
		exp2, err2 = Calc(expression[i+1:])
	}
	if err1 != nil {
		return 0.0, err1
	}
	if err2 != nil {
		return 0.0, err2
	}
	return do_exp(exp1, exp2, r1), nil
}

func Calc(expression string) (float64, error) {
	// работа со скобками
	expression = strings.ReplaceAll(expression, " ", "")
	ans := 0.0
	for i, r1 := range expression {
		if string(r1) == "(" {
			for j := range len(expression) {
				j = len(expression)-1-j
				if string([]rune(expression)[j]) == ")" {
					no_brakets, err := Calc(expression[i+1:j]) // 121, nil
					if err != nil {
						return 0.0, err
					}
					expression = expression[:i] + strconv.FormatFloat(no_brakets, 'f', -1, 64) + expression[j+1:]
					
					return Calc(expression)
				}
			}
			return 0.0, errors.New("closing parenthesis not found")
		}
		if string(r1) == "*" || string(r1) == "/" {
			return do_checks(expression, string(r1), i)
		}
		if string(r1) == "+" || string(r1) == "-" {
			return do_checks(expression, string(r1), i)
		}
	}
	ans, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return 0.0, err
	}
	return ans, nil
}

func main() {
	fmt.Println(Calc("2*2+2"))
}
