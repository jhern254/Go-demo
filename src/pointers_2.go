package main

import "fmt"

func main() {
	// TODO: Put all of this into a function 
	i, j := 42, 2701

	p := &i         // points to i, ref. value to i (int)
//	p := *i         // deref. pointer to i, gives value
//  z := *i // z := 42, var z int = 42
//	fmt.Println(z)
//   p := *i		// Go throws an error here, lang. is safeguarded against this, memory access
	// Pointers can only show value, but not copy over (assign)

	fmt.Println(p)  // 0xc000126000, points to memory address

	if false {
	    fmt.Println(*p) // 
	    *p = 21         // 
	    fmt.Println(i)  // 

	    p = &j         // 
	    *p = *p / 37   // 
	    fmt.Println(j) //
	}
}

/* 
Every langauge. Having a try bool loop is useful in literally every
language if you need to keep running code. 
*/
