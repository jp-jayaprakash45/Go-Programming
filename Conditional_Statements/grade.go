package main

import "fmt"

func main(){

	var marks int

		fmt.Print("enter your marks :")
		fmt.Scanf("%d",&marks)

		if marks >= 85 {
			fmt.Println("Grade : A ")
		}else if marks >= 70 {
			fmt.Println("Grade : B ")
		}else if marks >= 50 {
			fmt.Println("Grade : C ")
		}else if marks >= 35 {
			fmt.Println("Grade : D ")
		}else {
			fmt.Println("FAIL....")
		}
}
