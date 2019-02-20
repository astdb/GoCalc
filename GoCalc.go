
package main

import (
	"fmt"
	"os"
	"log"
	"regexp"
	"strings"
)

func main () {
	// check input exists
	if len(os.Args) != 2 {
		log.Fatal("Usage: $> go run GoCalc.go \"<expression>\"")
	}

	// read expression to be evaluated
	exp := strings.TrimSpace(os.Args[1])

	// check for validity
	// regex: one or more digits followed by zero or more of (one add/sub/div/mult signs followed by another digit)
	match, err := regexp.MatchString("^[0-9]\\s*([\\+|-|/|*]\\s*[0-9])*$", exp)
	if err != nil {
		log.Fatal(err)
	}

	// stacks for operators and operands
	operatorStack := []string{}
	operandStack := []string{}

	if match {
		// valid arithmetic expression - eval
		fmt.Println("Valid expression")

		// tokenize - operators and operands can be interspersed by arbitrary amounts of whitespace
		exp_runes := []rune(exp)
		exp_tokens := []string{}

		curTerm := ""
		index := 0
		for len(operandStack) > 1 || index == 0 {
			// put operators into operator stack
			// put operands into operand stack
			// every time an operand is pushed, check if the last operator is priority
			// if yes, operate on that operator and last popped op from operand stack
			// and push back onto operand stack (and pop priority operator from operator stack)
			// else, push until end and operate in order
			for i < len(exp_runes) {
				if isOperator(exp_runes[i]) {
					curTerm = strings.TrimSpace(curTerm)

					if isPriorityOperator(exp_runes[i]) {
						// pop last operand and operate
						prevOpd = operandStack[len(operandStack)-1]
						operandStack = operandStack[:len(operandStack)-1]
						operandStack = append

					} else {
						// push onto operand stack
						operandStack = append(operandStack, curTerm)
						curTerm = ""
					}
				} else {
					// keep building current term
					curTerm = fmt.Sprintf("%s%s", curTerm, string(exp_runes[i]))
				}
			}

			if index > len(exp_runes)-1 {
				break
			}


		}

		// for i := 0; i < len(exp_runes); i++ {
		// 	if isOperator(exp_runes[i]) {
		// 		if isPriorityOperator {
		// 			// operate on last two operands in operand stack
		// 			op1 := operandStack

		// 		} else {
		// 			// add to operator stack 
		// 			operatorStack = append(operatorStack, exp_runes[i])
		// 		}

		// 	} else {
		// 		// operand: add to operand stack
		// 		operandStack = append(operandStack, exp_runes[i])
		// 	}
		// }

		// operandStack := []string{}

	} else {
		log.Fatal("Invalid expression.")
	}
}

func isOperator(x rune) bool {
	if x == '+' || x == '-' || x == '/' || x == '*' {
		return true
	}

	return false
}

func isPriorityOperator(x rune) bool {
	if x == '/' || x == '*' {
		return true
	}

	return false
}
