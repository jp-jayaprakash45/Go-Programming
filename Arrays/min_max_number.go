package main

import "fmt"

func main(){

   var arr [5]int

   for i := 0;i<len(arr);i++{

              fmt.Printf("Enter arr[%d] :",i)
              fmt.Scanf("%d",&arr[i])
     }
    var max,min int

         min = arr[0]

    for i := 0;i<len(arr);i++{
             if  max < arr[i]{
                 max = arr[i]
              }
             if  min > arr[i]{
                 min = arr[i]
              }
     }
   fmt.Printf("Maximum = %d  Min = %d \n",max,min)


}
