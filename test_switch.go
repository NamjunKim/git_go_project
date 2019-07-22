package main

import "fmt"

func main() {

//	normal_number_case := 1
//	string_number_case := '9'
//	lower_string_case := 'a'
	upper_string_case := 'B'

	c := upper_string_case

	switch {
		//Go lang doesn't need parenthesize to express condition. So as is for 'if'
		//고는 조건문을 명시할 때 괄호를 필요로 하지 않는다. if문도 마찬가지.
		case 0 <= c  && c <= 9:
			fmt.Printf("%d is a number.", c)
	
		case '0' <= c && c<= '9':
			fmt.Printf("%c is a string number", c) 

		case 'a' <=c && c<= 'z':
			fmt.Printf("%c is a lower string", c)
			
		case 'A' <=c && c<= 'Z':
			fmt.Printf("%c is a upper string", c)

	}

}
