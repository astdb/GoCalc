
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

	// .. and check it's valid
	// regex: one or more digits followed by zero or more of (one add/sub/div/mult signs followed by another digit)
	match, err := regexp.MatchString("^[0-9]([\\+|-|/|*][0-9])*$", exp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(match)
}

