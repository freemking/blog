package main

import (
	"fmt"
	"func/params"
	"time"
)

func main() {
	exts := []params.CartOption{
		params.WithCartType1("buyNow"),
		params.WithTTL1(10 * time.Second),
	}
	foo := params.NewCart1("dayu", 888, exts...)
	fmt.Println(foo)
}
