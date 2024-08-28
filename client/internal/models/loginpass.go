package models

import "reflect"

type LoginPass struct {
	Login string
	Pass  string
}

var _ ValueInterface = (*LoginPass)(nil)

func (l *LoginPass) GetType() reflect.Type {
	return reflect.TypeOf(*l)
}
