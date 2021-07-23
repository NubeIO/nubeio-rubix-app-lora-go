package main

import "flag"
import "fmt"

func main() {


	wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	//fmt.Println("numb:", *numbPtr)
	//fmt.Println("fork:", *boolPtr)
	//fmt.Println("svar:", svar)
	//fmt.Println("tail:", flag.Args())
}
