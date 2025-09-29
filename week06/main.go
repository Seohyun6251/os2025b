package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var 64f float64
	// fmt.Println(64f,reflact.type(64f))

	totalPrice := 1000

	fmt.Println(totalPrice, reflect.TypeOf(totalPrice))
}
