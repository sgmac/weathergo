package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgmac/weathergo"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [city]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(-1)
	}

	result, err := weathergo.ByCity(args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
	fmt.Println(result)

}
