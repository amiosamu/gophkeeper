package models

import "reflect"

type Text struct {
	Text string
}

var _ ValueInterface = (*Text)(nil)

func (t *Text) GetType() reflect.Type {
	return reflect.TypeOf(*t)
}
