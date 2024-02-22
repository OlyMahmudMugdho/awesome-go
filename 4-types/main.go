package main

import (
	"fmt"
	"reflect"
)

func main(){
	var age  int = 23;
	var user string = "Mugdho"

	const pi float32 = 3.14

	fmt.Printf("The type of age is : %T \n", age)
	fmt.Printf("The type of user is : %T \n", user)

	fmt.Println("The type of pi is : ", reflect.TypeOf(pi))
}