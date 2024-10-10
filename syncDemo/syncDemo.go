package syncdemo

import (
	"fmt"
)

func IntegerGrnrator() chan int {
	ch := make(chan int)
	go func() {
		var i int
		for {
			ch <- i
			i++
		}
	}()

	go func() {
		var i int
		for {
			ch <- i
			i++
		}
	}()
	return ch
}

func PrintOutIntegerGrnrator() {
	ch := IntegerGrnrator()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
