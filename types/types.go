package types

import (
	"golang.org/x/crypto/bcrypt"
)

// Password: ******
type Password string

// MarshalJSON: json
func (p Password) MarshalJSON() (text []byte, err error) {
	return []byte(`"*******"`), nil
}

// Bytes
func (p Password) Bytes() []byte {
	return []byte(p)
}

// Generate
func (p Password) Generate() (Password, error) {
	var pw Password
	psd, err := bcrypt.GenerateFromPassword(p.Bytes(), 13)
	if err != nil {
		return pw, err
	}
	return Password(psd), err
}

// Compare
func (p Password) Compare(passwd []byte) error {
	return bcrypt.CompareHashAndPassword(p.Bytes(), passwd)
}
