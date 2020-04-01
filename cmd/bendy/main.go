package main

import (
	"flag"
	"os"
)

var file = flag.String("file", "", "source file")

func main() {
	flag.Parse()

	args := flag.Args()

	if *file == "" {
		if len(args) == 0 {
			panic("provide a file")
		} else {
			*file = args[0]
			args = args[1:]
		}
	}

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

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
