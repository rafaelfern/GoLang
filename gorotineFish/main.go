// Originally from: https://gobyexample.com/channels
// [Reprinted under license](https://creativecommons.org/licenses/by/3.0/).
// Poetic changes were made.
package main

import (
	"fmt"
	"time"
)

func printSlowly(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(time.Now(), i, s)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	printSlowly("directly functioning", 3)

	go printSlowly("red fish goroutine", 3)
	go printSlowly("blue fish goroutine", 3)

	go func(ss string, nn int) {
		for i := 0; i < nn; i++ {
			fmt.Println(i, ss)
			time.Sleep(150 * time.Millisecond)
		}
	}("anony fish goroutine", 3)

	var input string
	fmt.Scanln(&input)
	fmt.Println("DONE.")

}
