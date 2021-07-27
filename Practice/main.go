package main

import (
	"fmt"

	"rsc.io/quote/v3"

	mypkg "github.com/tkxkd0159/go-chain/hello"
	"github.com/tkxkd0159/go-chain/hello/nums"
)

func main() {
    fmt.Println(quote.GoV3())
    fmt.Println(mypkg.Hello())
    fmt.Println(nums.IsPrime(3))
}
