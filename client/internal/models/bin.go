package models

import "reflect"

type Bin struct {
	Bin []byte
}

var _ ValueInterface = (*Bin)(nil)

func (b *Bin) GetType() reflect.Type {
	return reflect.TypeOf(*b)
}
