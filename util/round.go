package util

import (
	"fmt"
	"math"

	"github.com/thesephist/ink/pkg/ink"
)

func RoundFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("round(num, [places])")

	if len(args) == 1 {
		args = append(args, ink.NumberValue(0))
	}
	if len(args) != 2 {
		return nil, usage
	}

	var num float64
	var places int

	if numv, ok := args[0].(ink.NumberValue); ok {
		num = float64(numv)
	} else {
		return nil, usage
	}

	if placesv, ok := args[1].(ink.NumberValue); ok {
		places = int(placesv)
	} else {
		return nil, usage
	}

	mult := math.Pow10(-places)

	num = math.Round(num / mult) * mult

	return ink.NumberValue(num), nil
}
