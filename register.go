package bendy

import (
	"fmt"

	"github.com/superloach/ink/pkg/ink"
)

func (b *Bendy) RegisterFn(ctx *ink.Context, args []ink.Value) (ink.Value, error) {
	usage := fmt.Errorf("register(event, fn)")

	if len(args) != 2 {
		return nil, usage
	}

	var name string
	var fn ink.Value

	if namev, ok := args[0].(ink.StringValue); ok {
		name = string(namev)
	} else {
		return nil, usage
	}

	if fnv, ok := args[1].(ink.FunctionValue); ok {
		fn = fnv
	} else if nfnv, ok := args[1].(ink.NativeFunctionValue); ok {
		fn = nfnv
	} else {
		return nil, usage
	}

	b.Funcs[name] = fn

	return nil, nil
}
