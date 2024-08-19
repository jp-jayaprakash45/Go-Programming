package main

import "fmt"

func main(){

	var input int

		fmt.Print("Enter any number from 1 - 7 :")
		fmt.Scanf("%d",&input)

		switch input {

			case 1 :
				fmt.Println("MONDAY")

			case 2 :
				fmt.Println("TUESDAY")

			case 3 :
				fmt.Println("WEDNESDAY")

			case 4 :
				fmt.Println("THURSDAY")

			case 5 :
				fmt.Println("FRIDAY")

			case 6 :
				fmt.Println("SATURDAY")

			case 7 :
				fmt.Println("SUNDAY")

			default :
				fmt.Println("Pick appropriate number between 1 - 7")

		}
}


