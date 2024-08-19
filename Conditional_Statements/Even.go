package main

import "fmt"

func main(){

	var input int 

		fmt.Print("enter any number : ")
		fmt.Scanf("%d",&input)

		fmt.Printf("printing even numbers till %d \n",input)
		for i:=0 ; i <= input ; i++  {
			if i %2 == 0  {
				 fmt.Println(i)
			}


		}
}
