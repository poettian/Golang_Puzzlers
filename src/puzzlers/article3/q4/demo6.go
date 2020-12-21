package main

import (
	"flag"
	"puzzle/src/puzzlers/article3/q4/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout, name)
}
