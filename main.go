package main

import (
	"fmt"
	"koda-b8-golang4/libs"
)

func main() {
	res := libs.SearchPerson([]string{"john", "jane", "json"}, "jane")
	fmt.Println(res)
}
