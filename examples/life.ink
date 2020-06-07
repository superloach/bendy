sz := size()

s := {
	cells: {}
}

set := (c, x, y, v) => (
	c.(y).(x) := v
	poke(x, y, v * 3 * (4 + (x + y) % 60))
)

get := (c, x, y) => (
	y + 1 :: {
		0 -> y := sz.h - 1
		(sz.h + 1) -> y := 0
	}
	x + 1 :: {
		0 -> x := sz.w - 1
		(sz.w + 1) -> x := 0
	}
	c.(y).(x) :: {
		() -> 0
		_ -> c.(y).(x)
	}
)

neighbors := (c, x, y) => (
	get(c, x - 1, y - 1) +
	get(c, x - 1, y) +
	get(c, x - 1, y + 1) +
	get(c, x, y - 1) +
	get(c, x, y + 1) +
	get(c, x + 1, y - 1) +
	get(c, x + 1, y) +
	get(c, x + 1, y + 1)
)

randcells := () => each(seq(0, sz.h), y => (
	s.cells.(y) := {}
	each(seq(0, sz.w), x => (
		set(s.cells, x, y, round(rand()))
	))
))

each(seq(0, sz.h), y => (
	s.cells.(y) := {}
	each(seq(0, sz.w), x => (
		set(s.cells, x, y, 0)
	))
))

set(s.cells, 0, 0, 1)
set(s.cells, 2, 0, 1)
set(s.cells, 1, 1, 1)
set(s.cells, 2, 1, 1)
set(s.cells, 1, 2, 1)

register('update', () => (
	c := {}
	each(seq(0, sz.h), y => (
		c.(y) := {}
		each(seq(0, sz.w), x => (
			n := neighbors(s.cells, x, y)
			get(s.cells, x, y) :: {
				1 -> n :: {
					2 -> set(c, x, y, 1)
					3 -> set(c, x, y, 1)
					_ -> set(c, x, y, 0)
				}
				0 -> n :: {
					3 -> set(c, x, y, 1)
					_ -> set(c, x, y, 0)
				}
			}
		))
	))
	s.cells := c
))

register('key', n => randcells())
