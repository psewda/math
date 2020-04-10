package main

import (
	"fmt"

	"github.com/rs/xid"
)

type emp struct {
	Name string
	name string
}

func main() {

	uid := xid.New().String()

	fmt.Println("Hello Math => " + uid)

	e := emp{
		Name: "N",
		name: "v",
	}

	fmt.Println(e)
}
