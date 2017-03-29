package main

import "time"
import "fmt"

func main() {
   switch time.Now().Weekday() {
      case time.Saturday, time.Sunday:
         fmt.Println("Weekend")
      default:
         fmt.Println("Weekday")
   }
}
