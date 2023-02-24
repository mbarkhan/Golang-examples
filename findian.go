package main

import (
	"fmt"
	 "strings"
)


func main() {
	var str string

	fmt.Printf("Enter a string : ")

	fmt.Scan(&str);
	if len(str)>2{
		ustr := strings.ToUpper(str)
		if strings.HasPrefix(ustr, "I") && strings.Contains(ustr, "A") && strings.HasSuffix(ustr, "N") {
				fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else{
		fmt.Println("Error in Input string")
	}
	}
