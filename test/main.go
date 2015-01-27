package main

import (
	"log"
	"os"

	"github.com/juju/errgo"
	"github.com/altipla-consulting/arb-reader/arbreader"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f, err := os.Open("test/test.arb")
	if err != nil {
		return errgo.Mask(err)
	}
	defer f.Close()

	messages, err := arbreader.Read(f)
	if err != nil {
		return errgo.Mask(err)
	}

	log.Printf("\n\t%+v\n\t%+v", messages[0], messages[1])

	return nil
}
