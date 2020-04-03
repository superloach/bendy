package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/superloach/bendy"
	"github.com/superloach/ink/pkg/ink"
)

var keys = map[string]ebiten.Key{
	"up":    ebiten.KeyUp,
	"down":  ebiten.KeyDown,
	"left":  ebiten.KeyLeft,
	"right": ebiten.KeyRight,
	"a":     ebiten.KeyZ,
	"b":     ebiten.KeyX,
	"c":     ebiten.KeyC,
	"menu":  ebiten.KeyEscape,
}

type game struct {
	bendy *bendy.Bendy
	disp  *ebiten.Image
}

func newGame() (*game, error) {
	g := &game{}

	disp, err := ebiten.NewImage(
		bendy.DispWidth,
		bendy.DispHeight,
		ebiten.FilterDefault,
	)
	if err != nil {
		return nil, err
	}
	g.disp = disp

	b := bendy.NewBendy()
	b.Poke = g.Poke
	g.bendy = b

	return g, nil
}

func (g *game) Layout(ow, oh int) (int, int) {
	return bendy.DispWidth, bendy.DispHeight
}

func (g *game) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if fn, ok := g.bendy.Funcs["update"]; ok {
		_, err := g.bendy.Ctx.EvalFunc(fn, false)
		if err != nil {
			return err
		}
	}

	for n, k := range keys {
		if fn, ok := g.bendy.Funcs["keydown"]; ok {
			if inpututil.IsKeyJustPressed(k) {
				_, err := g.bendy.Ctx.EvalFunc(
					fn, false,
					ink.StringValue(n),
				)
				if err != nil {
					return err
				}
			}
		}
		if fn, ok := g.bendy.Funcs["keyup"]; ok {
			if inpututil.IsKeyJustReleased(k) {
				_, err := g.bendy.Ctx.EvalFunc(
					fn, false,
					ink.StringValue(n),
				)
				if err != nil {
					return err
				}
			}
		}
	}

	screen.DrawImage(g.disp, &ebiten.DrawImageOptions{})

	return nil
}

func (g *game) Poke(x, y, color int) {
	g.disp.Set(x, y, bendy.DispColors[color])
}

func (g *game) Run() error {
	return ebiten.RunGame(g)
}
