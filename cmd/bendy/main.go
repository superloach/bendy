package main

import (
	"flag"
	"io"
	"os"

	"github.com/gobuffalo/packr/v2"
)

var (
	file = flag.String("file", "", "source code")

	width = flag.Int("width", 128, "screen width")
	height = flag.Int("height", 96, "screen height")
)

var box = packr.New("builtin game", ".")

func main() {
	flag.Parse()
	args := flag.Args()

	var i io.Reader

	if box.Has("builtin.ink") {
		f, err := box.Open("builtin.ink")
		if err != nil {
			panic(err)
		}
		i = f
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
				i = f
			}
		}
	}

	game, err := newGame(*width, *height)
	if err != nil {
		panic(err)
	}

	err = game.bendy.Load(i)
	if err != nil {
		panic(err)
	}

	err = game.Run()
	if err != nil {
		panic(err)
	}
}
