package main

import (
	"flag"
	"fmt"
)

func main() {
	add := flag.Bool("add", false, "a bool")
	//list := flag.String("list", "foo", "a string")
	//delete := flag.String("delete", "foo", "a string")
	//summary := flag.String("summary", "foo", "a string")
	amount := flag.Int("amount", 0, "a integer")

	//id := flag.String("id  ", "foo", "a string")
	description := flag.String("description", "foo", "a string")
	flag.Parse()
	if *add {
		fmt.Println(*description, *amount)
	}

	//fmt.Println(*add, *list, *delete, *summary, *amount, *id)
	//Hello

}
