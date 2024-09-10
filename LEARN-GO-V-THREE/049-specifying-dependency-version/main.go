package main

import (
	
	"fmt"
	"github.com/anuvara2000/puppy"
)
func main() {

	puppy.Form11()
	puppy.Form12()
	puppy.Form13()

      s1 :=puppy.Bark()
	  s2 :=puppy.Barks()
	  fmt.Println(s1)
	  fmt.Println(s2)
       
	  s3:=puppy.BigBark()
	  s4:=puppy.BigBarks()
      fmt.Println(s3)
	  fmt.Println(s4)

	  // also like this
	  fmt.Println(puppy.Bark())
	  fmt.Println(puppy.Barks())
}