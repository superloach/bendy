package util

import (
	"fmt"

	"github.com/superloach/ink/pkg/ink"
)

func LogFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	vs := make([]interface{}, len(args))
	for i, arg := range args {
		vs[i] = arg
	}
	fmt.Println(vs...)

	return nil, nil
}
