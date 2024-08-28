package models

import "reflect"

type Card struct {
	Owner  string
	Number string
	Date   string
	CVV    string
}

var _ ValueInterface = (*Card)(nil)

func (c *Card) GetType() reflect.Type {
	return reflect.TypeOf(*c)
}
