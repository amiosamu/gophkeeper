package models

type Action byte

const (
	Add Action = iota
	Modify
	Delete
)

type Type byte

const (
	TloginPass Type = iota
	Ttext
	Tbin
	Tcard
)

type Record struct {
	Id         int64
	Action     Action
	RecordType Type
	Value      []byte
	Meta       string
}

func NewRecord(
	id int64,
	action Action,
	recordType Type,
	key []byte,
	v ValueInterface,
	meta string,
) *Record {
	magicKey := make([]byte, 0, 16)
	magicKey = append(magicKey, key[:16]...)

	return &Record{
		Id:         id,
		Action:     action,
		RecordType: recordType,
		Value:      EncryptAES(magicKey, v),
		Meta:       meta,
	}
}

func (r *Record) DecryptValue(key []byte) ValueInterface {
	magicKey := make([]byte, 0, 16)
	magicKey = append(magicKey, key[:16]...)

	return DecryptAES(magicKey, r.Value)
}
