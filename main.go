package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func usage() {
	fmt.Print(`Usage: appendnl

Will add a newline (\n) to stdin.

EXAMPLE:
	$ echo -n "foobar" | appendnl 
`)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() > 1 {
		usage()
		os.Exit(1)
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input := string(data)
	fmt.Printf("%s\n", input)
}
