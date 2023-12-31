package main

import (
	"log"

	"github.com/pkg/errors"

	origerrors "errors"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := baz()
	if err != nil {
		log.Printf("%+v", errors.FirstStack(err))
	}
}

func libfunc() error {
	return origerrors.New("an error occurred in foo")
}

func bar() error {
	return errors.Wrap(libfunc(), "an error occurred in bar")
}

func baz() error {
	return errors.Wrap(bar(), "an error occurred in baz")
}
