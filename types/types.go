package types

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) MarshalJSON() (text []byte, err error) {
	return []byte(`"*******"`), nil
}

func (p Password) Bytes() []byte {
	return []byte(p)
}

func (p Password) Generate() (Password, error) {
	var pw Password
	psd, err := bcrypt.GenerateFromPassword(p.Bytes(), 13)
	if err != nil {
		return pw, err
	}
	return Password(psd), err
}

func (p Password) Compare(passwd []byte) error {
	return bcrypt.CompareHashAndPassword(p.Bytes(), passwd)
}
