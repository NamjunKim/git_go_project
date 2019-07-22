package main

import "fmt"

func main(){

	sum, i := 0, 0
	for i < 10 {
	
		
		// impossible using prefix calculation
		// sum+=i
		// ++i

		// impossible using below way
		// sum+=i++

		sum+=i

		i++

	}

	fmt.Println("Result : ",  sum);

}
