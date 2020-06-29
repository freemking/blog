package main

import (
	"fmt"
	"func/params"
	"time"
)

func main() {
	ext := params.NewExt()
	ext.WithCartType("buyNow")
	ext.WithTTL(10 * time.Second)

	fmt.Println(params.NewCart("dayu", 888, ext))
}
