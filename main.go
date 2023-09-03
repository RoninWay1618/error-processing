package main

import "fmt"

func main() {
	var i interface{} = "hello"

	defer func() {
		if f := recover(); f != nil {
			fmt.Println("Восстановаление из паники:", f)
		}
	}()

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
