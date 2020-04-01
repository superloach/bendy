` FLAPPY - jumping bird game for bendy `

` vars `
state      := { page: 0 }
width      := 128
height     := 96
bgColor    := 0  ` 000000, black `
titleColor := 63 ` 111111, white `
birbColor  := 56 ` 111000, orange `
pipeColor  := 12 ` 001100, lime `
birbX      := width / 6
birbGrav   := height / 192
birbJump   := height / 12
pipeGap    := birbJump * 2.5
pipeY      := () => (height / 6) + rand() * (2 * height / 3 - pipeGap)
count      := 3
titleText  := 'FLAPPY
for bendy

press any key!'

register('keydown', n => state.page :: {
	` title screen: -> playing `
	0 -> (
		clear(bgColor)
		state.birbY := height / 2
		state.score := 0

		state.pipes := {}
		each(seq(0, count + 1), n => (
			state.pipes.(n) := {}
			state.pipes.(n).x := (width / count) * (n + 1)
			state.pipes.(n).y := pipeY()
		))

		state.page := 1
	)

	` playing: jump `
	1 -> (
		poke(birbX, height - 1 - state.birbY, bgColor)
		state.birbY := state.birbY + birbJump
		poke(birbX, height - 1 - state.birbY, birbColor)
	)
})

pipeDraw := (x, y, gap, color) => (
	line(x, height, x, height - y, color)
	line(x, height - (y + gap), x, 0, color)
)

score := () => (
	each(seq(0, width), x => line(x, height - 8, x, height, 0))
	text(state.score, 1, height - 8)
)

register('update', () => state.page :: {
	0 -> ( ` title screen `
		text(titleText, 1, 1)
	)

	1 -> ( ` playing `
		` score `
		score()

		` clear birb `
		line(birbX, 0, birbX, height, bgColor)

		` draw pipes `
		each(seq(0, count + 1), n => (
			pipeDraw(state.pipes.(n).x, state.pipes.(n).y, pipeGap, bgColor)
			state.pipes.(n).x < 1 :: { true -> (
				state.pipes.(n).x := width + (width / count)
				state.pipes.(n).y := pipeY()
			) }
			state.pipes.(n).x := state.pipes.(n).x - 1
			pipeDraw(state.pipes.(n).x, state.pipes.(n).y, pipeGap, pipeColor)

			` collision `
			round(state.pipes.(n).x) = round(birbX) :: { true -> (
				state.birbY < state.pipes.(n).y :: {
					true -> ( state.page := 0 )
					false -> state.birbY > state.pipes.(n).y + pipeGap :: {
						true -> ( state.page := 0 )
						false -> (
							state.score := state.score + 1
							score()
						)
					}
				}
			) }
		))

		` draw birb `
		state.birbY > height :: { true -> ( state.birbY := height ) }
		state.birbY < 1 :: { true -> ( state.page := 0 ) }
		state.birbY := state.birbY - birbGrav
		poke(birbX, height - state.birbY, birbColor)
	)
})
