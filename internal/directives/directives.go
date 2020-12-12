package directives

import (
	"context"
	"fmt"

	"github.com/logrusorgru/gqlif/internal/types"

	"github.com/99designs/gqlgen/graphql"
)

type ValueTypeAllow struct {
	ValueType types.ValueType `json:"valueType"`
	Allow     bool            `json:"allow"`
}

func AllowValueType(ctx context.Context, obj interface{}, next graphql.Resolver,
	field string, allow []*ValueTypeAllow) (
	res interface{}, err error) {

	println("AllowValueType has called with:", fmt.Sprintf("%T", obj))

	return next(ctx)
}
