package main

import (
	"flag"
	"fmt"
)

func main() {
	add := flag.String("add", "foo", "a string")
	list := flag.String("list", "foo", "a string")
	delete := flag.String("delete", "foo", "a string")
	summary := flag.String("summary", "foo", "a string")
	amount := flag.String("amount ", "foo", "a string")
	id := flag.String("id  ", "foo", "a string")

	flag.Parse()
	fmt.Println(*add, *list, *delete, *summary, *amount, *id)

}
