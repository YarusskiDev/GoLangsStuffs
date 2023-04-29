package main

import (
	"fmt"
	"reflect"
)

func main() {

	slice := make([]int, 10)
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	slice = append(slice, 1)
	fmt.Println(slice)
	slice[0] = 1001
	fmt.Println(slice)
	newSlice := append(slice, 11)
	fmt.Println(&newSlice)
	sliceSemReferencia := make([]string, 10)
	sliceSemReferencia[1] = "Cristiano"
	fmt.Println(sliceSemReferencia)

}
