sz := size()

state := {
	x: 64,
	y: 64,
	c: 42,
}

each(seq(0, 64), n =>
	each(seq(0, 3), i =>
		each(seq(0, 3), j =>
			poke((n % sz.w) * 2 + i, (n / sz.w) * 2 + j, n)
		)
	)
)

t := 'Hello, World!

1234567890ABCDEFGHIJK
LMNOPQRSTUVWXYZabcdef
ghijklmnopqrstuvwxyz`
~!@#$%^&*()-_=+[{]}\\|
;:\'",<.>/?'

mksprite('foo', 3, '.... ....')

register('update', () => (
	clear()
	sprite('foo', state.x, state.y, state.c)
	text(t, 1, 1, state.c)
))

register('key', n => (
	n :: {
		'up' -> state.y := state.y - 1
		'down' -> state.y := state.y + 1
		'left' -> state.x := state.x - 1
		'right' -> state.x := state.x + 1
		'a' -> state.c := state.c + 1
		'b' -> state.c := state.c - 1
	}
))
