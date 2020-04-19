package bendy

import (
	"io"

	"github.com/superloach/bendy/util"
	"github.com/superloach/ink/pkg/ink"
)

type Bendy struct {
	Width   int
	Height  int
	Funcs   map[string]ink.Value
	Poke    func(x, y, color int)
	Ctx     *ink.Context
	Sprites map[string]*Sprite
}

func NewBendy(w, h int) *Bendy {
	b := &Bendy{}

	b.Width = w
	b.Height = h

	b.Funcs = make(map[string]ink.Value)
	b.Sprites = make(map[string]*Sprite)

	engine := &ink.Engine{}
	engine.Permissions.Read = false
	engine.Permissions.Write = false
	engine.Permissions.Net = false
	engine.Permissions.Exec = false

	ctx := engine.CreateContext()
	ctx.LoadFunc("size", b.SizeFn)
	ctx.LoadFunc("register", b.RegisterFn)
	ctx.LoadFunc("log", util.LogFn)
	ctx.LoadFunc("round", util.RoundFn)
	ctx.LoadFunc("each", util.EachFn)
	ctx.LoadFunc("seq", util.SeqFn)
	ctx.LoadFunc("poke", b.PokeFn)
	ctx.LoadFunc("clear", b.ClearFn)
	ctx.LoadFunc("line", b.LineFn)
	ctx.LoadFunc("text", b.TextFn)
	ctx.LoadFunc("mksprite", b.MkSpriteFn)
	ctx.LoadFunc("sprite", b.SpriteFn)
	b.Ctx = ctx

	return b
}

func (b *Bendy) Load(i io.Reader) error {
	_, err := b.Ctx.Exec(i)
	return err
}
