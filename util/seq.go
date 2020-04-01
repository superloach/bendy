package util

import (
	"fmt"
	"strconv"

	"github.com/thesephist/ink/pkg/ink"
)

func SeqFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("seq(min, max, [inc])")

	if len(args) == 2 {
		args = append(args, ink.NumberValue(1))
	}
	if len(args) != 3 {
		return nil, usage
	}

	var min float64
	var max float64
	var inc float64

	if minv, ok := args[0].(ink.NumberValue); ok {
		min = float64(minv)
	} else {
		return nil, usage
	}

	if maxv, ok := args[1].(ink.NumberValue); ok {
		max = float64(maxv)
	} else {
		return nil, usage
	}

	if incv, ok := args[2].(ink.NumberValue); ok {
		inc = float64(incv)
	} else {
		return nil, usage
	}

	seq := ink.CompositeValue{}
	n := min
	i := 0
	for n < max {
		n += inc
		i++
		seq[strconv.Itoa(i)] = ink.NumberValue(n)
	}

	return seq, nil
}
