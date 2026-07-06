package main

import (
	"fmt"
	"koda-b8-golang4/libs"
)

func main() {
	name := ""
	fmt.Print("Masukan nama yang ingin dicari: ")
	fmt.Scanf("%s", &name)
	res := libs.SearchPerson([]string{"John", "Jane", "Json"}, name)
	fmt.Println(res)
}
