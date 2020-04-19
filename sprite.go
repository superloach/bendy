package bendy

import (
	"fmt"

	"github.com/superloach/ink/pkg/ink"
)

type Sprite struct {
	Data string
	Wrap int
}

func (b *Bendy) sprite(s *Sprite, x, y, color int) {
	for i := 0; i < len(s.Data); i++ {
		if s.Data[i] == ' ' {
			continue
		}
		b.poke(x+i%s.Wrap, y+i/s.Wrap, color)
	}
}

func (b *Bendy) MkSpriteFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("mksprite(name, wrap, data)")

	if len(args) != 3 {
		return nil, usage
	}

	var name string
	var wrap int
	var data string

	if namev, ok := args[0].(ink.StringValue); ok {
		name = string(namev)
	} else {
		return nil, usage
	}

	if wrapv, ok := args[1].(ink.NumberValue); ok {
		wrap = int(wrapv)
	} else {
		return nil, usage
	}

	if datav, ok := args[2].(ink.StringValue); ok {
		data = string(datav)
	} else {
		return nil, usage
	}

	ndata := []rune{}
	for _, r := range data {
		if r != '\n' {
			ndata = append(ndata, r)
		}
	}

	b.Sprites[name] = &Sprite{
		Data: string(ndata),
		Wrap: wrap,
	}

	return nil, nil
}

func (b *Bendy) SpriteFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("sprite(name, x, y, [color])")

	if len(args) == 3 {
		args = append(args, ink.NumberValue(63))
	}
	if len(args) != 4 {
		return nil, usage
	}

	var name string
	var x, y int
	var color int

	if namev, ok := args[0].(ink.StringValue); ok {
		name = string(namev)
	} else {
		return nil, usage
	}

	if xv, ok := args[1].(ink.NumberValue); ok {
		x = int(xv)
	} else {
		return nil, usage
	}

	if yv, ok := args[2].(ink.NumberValue); ok {
		y = int(yv)
	} else {
		return nil, usage
	}

	if colorv, ok := args[3].(ink.NumberValue); ok {
		color = int(colorv)
	} else {
		return nil, usage
	}

	s, ok := b.Sprites[name]
	if !ok {
		return nil, fmt.Errorf("sprite %s not found", name)
	}
	b.sprite(s, x, y, color)

	return nil, nil
}
