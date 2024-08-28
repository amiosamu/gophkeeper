package models

import (
	"bytes"
	"crypto/aes"
	"encoding/gob"
	"log"
	"reflect"
)

type ValueInterface interface {
	GetType() reflect.Type
}

func EncryptAES(key []byte, v ValueInterface) []byte {
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(v)

	if err != nil {
		log.Fatal(err)
	}

	out := make([]byte, len(buf.Bytes()))

	c.Encrypt(out, buf.Bytes())
	return out
}

func DecryptAES(key []byte, cipher []byte) ValueInterface {
	var v ValueInterface

	c, err := aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	out := make([]byte, len(cipher))
	c.Decrypt(out, cipher)

	var buf bytes.Buffer

	n, err := buf.Write(out)
	if err != nil {
		log.Fatal(err)
	}
	if n != len(out) {
		log.Fatal("Incorrect length of decrypted message ")
	}

	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&v)

	if err != nil {
		log.Fatal(err)
	}

	return v
}
