package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func try(word, post, pre string) (x, y string) {
	x = word + " " + post
	y = pre + " " + word
	return 
}

func main() {
	fmt.Println(split(17))
	a, b := try("hello", "world", "say")
	fmt.Println(a, "\n", b)
}

