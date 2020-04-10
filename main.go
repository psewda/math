package main

import (
	"fmt"

	"github.com/rs/xid"
)

func main() {

	uid := xid.New().String()

	fmt.Println("Hello Math => " + uid)
}
