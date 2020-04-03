package main

import (
	"os"
	"io"
	"flag"

	"github.com/gobuffalo/packr/v2"
)

var file = flag.String("file", "", "source code")

var box = packr.New("builtin game", ".")

func main() {
	flag.Parse()
	args := flag.Args()

	var f io.Reader
	var err error

	if box.Has("builtin.ink") {
		f, err = box.Open("builtin.ink")
		if err != nil {
			panic(err)
		}
	} else {
		if *file == "" {
			if len(args) < 1 {
				println("please provide a -file flag")
				flag.Usage()
				return
			} else {
				*file = args[0]
				args = args[1:]

				f, err := os.Open(*file)
				if err != nil {
					panic(err)
				}
				defer f.Close()
			}
		}
	}

	game, err := newGame()
	if err != nil {
		panic(err)
	}

	err = game.bendy.Load(f)
	if err != nil {
		panic(err)
	}

	err = game.Run()
	if err != nil {
		panic(err)
	}
}
