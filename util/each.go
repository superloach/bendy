package util

import (
	"fmt"
	"strconv"

	"github.com/superloach/ink/pkg/ink"
)

func EachFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("each(list, fn)")

	if len(args) != 2 {
		return nil, usage
	}

	var list []ink.Value
	var fn ink.Value

	if listv, ok := args[0].(ink.CompositeValue); ok {
		list = make([]ink.Value, len(listv))

		for k, v := range listv {
			i, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}

			list[i - 1] = v
		}
	}

	if fnv, ok := args[1].(ink.Value); ok {
		fn = fnv
	}

	var vals = ink.CompositeValue{}

	for i, item := range list {
		val, err := ctx.EvalFunc(fn, false, item)
		if err != nil {
			return vals, err
		}

		vals[strconv.Itoa(i + 1)] = val
	}

	return vals, nil
}
