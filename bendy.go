package bendy

import (
	"io"

	"github.com/superloach/ink/pkg/ink"
	"github.com/superloach/bendy/util"
)

type Bendy struct {
	Funcs map[string]ink.Value
	Poke  func(x, y, color int)
	Ctx   *ink.Context
}

func NewBendy() *Bendy {
	b := &Bendy{}

	b.Funcs = make(map[string]ink.Value)

	engine := &ink.Engine{}
	engine.Permissions.Read = false
	engine.Permissions.Write = false
	engine.Permissions.Net = false
	engine.Permissions.Exec = false

	ctx := engine.CreateContext()
	ctx.LoadFunc("register", b.RegisterFn)
	ctx.LoadFunc("log", util.LogFn)
	ctx.LoadFunc("round", util.RoundFn)
	ctx.LoadFunc("each", util.EachFn)
	ctx.LoadFunc("seq", util.SeqFn)
	ctx.LoadFunc("poke", b.PokeFn)
	ctx.LoadFunc("clear", b.ClearFn)
	ctx.LoadFunc("line", b.LineFn)
	ctx.LoadFunc("text", b.TextFn)
	b.Ctx = ctx

	return b
}

func (b *Bendy) Load(i io.Reader) error {
	_, err := b.Ctx.Exec(i)
	return err
}
