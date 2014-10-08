package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgmac/weathergo"
)

func main() {
	args := os.Args
	ch := make(chan interface{})

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [city]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(-1)
	}

	for _, city := range args[1:] {
		go func() {
			result, err := weathergo.ByCity(city)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(-1)
			}
			ch <- result
		}()
		fmt.Println(<-ch)
	}
}
