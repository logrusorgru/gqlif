package types

import (
	"fmt"
	"io"
	"strconv"
)

type Value struct {
	Name      string    `json:"name"`
	ValueType ValueType `json:"valueType"`
}

type ValueType string

const (
	ValueTypeOne ValueType = "ONE"
	ValueTypeTwo ValueType = "TWO"
)

var AllValueType = []ValueType{
	ValueTypeOne,
	ValueTypeTwo,
}

func (e ValueType) IsValid() bool {
	switch e {
	case ValueTypeOne, ValueTypeTwo:
		return true
	}
	return false
}

func (e ValueType) String() string {
	return string(e)
}

func (e *ValueType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ValueType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ValueType", str)
	}
	return nil
}

func (e ValueType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
