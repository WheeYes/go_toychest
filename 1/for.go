package main

import "fmt"

func main() {
   i :=1

   // just like while loop
   for i <=3 {
      fmt.Println(i)
      i = i+1
   }
   
   // traditional for loop
   for j :=7; j <=9; j++ {
      fmt.Println(j)
   }
   
   // while 1 loop
   for {
      fmt.Println("forever loop")
      // break now
      break      
   } 
}
