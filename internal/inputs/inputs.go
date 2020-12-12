package inputs

import (
	"github.com/logrusorgru/gqlif/internal/types"
)

type InputValue struct {
	Name      string          `json:"name"`
	ValueType types.ValueType `json:"valueType"`
}
