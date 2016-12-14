package types

type Password string

func (p Password) MarshalJSON() (text []byte, err error) {
	return []byte(`"*******"`), nil
}
