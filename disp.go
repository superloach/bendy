package bendy

import (
	"fmt"
	"image/color"

	"github.com/superloach/ink/pkg/ink"
)

const DispNumColors = 1 << 6

var DispColors [DispNumColors]color.RGBA64

func init() {
	for i := 0; i < DispNumColors; i++ {
		q := uint32(i)
		r := uint16(((q & (0b11 << 4)) >> 4) * 0x5555)
		g := uint16(((q & (0b11 << 2)) >> 2) * 0x5555)
		b := uint16(((q & (0b11 << 0)) >> 0) * 0x5555)
		DispColors[i] = color.RGBA64{r, g, b, 0xFFFF}
	}
}

func (b *Bendy) poke(x, y, color int) {
	if x < 0 || x >= b.Width {
		return
	}

	if y < 0 || y >= b.Height {
		return
	}

	if b.Poke != nil {
		b.Poke(x, y, color)
	}
}

func (b *Bendy) SizeFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("size()")

	if len(args) != 0 {
		return nil, usage
	}

	return ink.CompositeValue{
		"w": ink.NumberValue(b.Width),
		"h": ink.NumberValue(b.Height),
	}, nil
}

func (b *Bendy) PokeFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("poke(x, y, color)")

	if len(args) != 3 {
		return nil, usage
	}

	var x, y, color int

	if xv, ok := args[0].(ink.NumberValue); ok {
		x = int(xv)
	} else {
		return nil, usage
	}

	if yv, ok := args[1].(ink.NumberValue); ok {
		y = int(yv)
	} else {
		return nil, usage
	}

	if colorv, ok := args[2].(ink.NumberValue); ok {
		color = int(colorv)
	} else {
		return nil, usage
	}

	if x < 0 || x >= b.Width {
		return nil, nil
	}

	if y < 0 || y >= b.Height {
		return nil, nil
	}

	b.poke(x, y, color)

	return nil, nil
}

func (b *Bendy) ClearFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("clear([color])")

	if len(args) == 0 {
		args = append(args, ink.NumberValue(0))
	}

	if len(args) != 1 {
		return nil, usage
	}

	var color int

	if colorv, ok := args[0].(ink.NumberValue); ok {
		color = int(colorv)
	} else {
		return nil, usage
	}

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			b.poke(x, y, color)
		}
	}

	return nil, nil
}

func (b *Bendy) LineFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("line(x1, y1, x2, y2, color)")

	if len(args) != 5 {
		return nil, usage
	}

	var x1, y1, x2, y2 float64
	var color int

	if x1v, ok := args[0].(ink.NumberValue); ok {
		x1 = float64(x1v)
	} else {
		return nil, usage
	}

	if y1v, ok := args[1].(ink.NumberValue); ok {
		y1 = float64(y1v)
	} else {
		return nil, usage
	}

	if x2v, ok := args[2].(ink.NumberValue); ok {
		x2 = float64(x2v)
	} else {
		return nil, usage
	}

	if y2v, ok := args[3].(ink.NumberValue); ok {
		y2 = float64(y2v)
	} else {
		return nil, usage
	}

	if colorv, ok := args[4].(ink.NumberValue); ok {
		color = int(colorv)
	} else {
		return nil, usage
	}

	if x1 > x2 {
		tmp := x2
		x2 = x1
		x1 = tmp
	}

	if y1 > y2 {
		tmp := y2
		y2 = y1
		y1 = tmp
	}

	dx := x2 - x1
	dy := y2 - y1

	if dx > dy {
		for x := 0.0; x < dx; x++ {
			y := dy * (x / dx)
			b.poke(int(x1+x), int(y1+y), color)
		}
	} else {
		for y := 0.0; y < dy; y++ {
			x := dx * (y / dy)
			b.poke(int(x1+x), int(y1+y), color)
		}
	}

	return nil, nil
}
