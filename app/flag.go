package app

type Flag struct {
	Name      string
	ShortName string
	Default   string
	Value     string
}

func newFlag(name, short, val string) *Flag {
	return &Flag{
		Name:      name,
		ShortName: short,
		Default:   val,
		Value:     val,
	}
}
