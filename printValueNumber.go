package main

import "fmt"

func addDigitOfnumber () {
    n := 1080
	var sum int
    var res int
   

   for n > 0 {
       res = n % 10
       sum = sum + res
       n = n/10
   }
   fmt.Println( "sum : ",sum)
}