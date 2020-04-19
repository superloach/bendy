state := {
	x: 64,
	y: 64,
	c: 0,
}

t := 'Hello, World!

1234567890ABCDEFGHIJK
LMNOPQRSTUVWXYZabcdef
ghijklmnopqrstuvwxyz`
~!@#$%^&*()-_=+[{]}\\|
;:\'",<.>/?'

mksprite('foo', 5, '
.....
.   .
.   .
.   .
.....
')

register('update', () => (
	sprite('foo', 64, 64, state.c)
	text(t, 1, 1, state.c)
))

register('key', () => (
	state.c := state.c + 1
))
