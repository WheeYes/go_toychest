package main

import "strings"
import "fmt"

func main() {
   fmt.Println("contains " , strings.Contains("test", "es"))
   fmt.Println("count " , strings.Count("hell", "hell"))
   fmt.Println("ToLower " , strings.ToLower("HELLO"))
}
