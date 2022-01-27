package myquote

import (

	"fmt"
	"rsc.io/quote"
)

func PrintQuote()  {

	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}

