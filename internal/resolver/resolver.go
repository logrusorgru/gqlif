package resolver

import (
	"context"
	"fmt"

	"github.com/logrusorgru/gqlif/internal/inputs"
	"github.com/logrusorgru/gqlif/internal/types"
)

type Resolver struct {
	reg map[string]*types.Value
}

func NewResolver() (r *Resolver) {
	r = new(Resolver)
	r.reg = make(map[string]*types.Value)
	return
}

func (r *Resolver) Set(ctx context.Context, name string,
	value inputs.InputValue) (val *types.Value, err error) {

	val = &types.Value{Name: value.Name, ValueType: value.ValueType}
	r.reg[name] = val

	return
}

func (r *Resolver) Get(ctx context.Context, name string) (
	val *types.Value, err error) {

	var ok bool
	if val, ok = r.reg[name]; !ok {
		return nil, fmt.Errorf("not found")
	}

	return
}
