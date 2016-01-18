package optional

type Bool struct {
	isSet bool
	value bool
}

func NewBool(value bool) Bool {
	return Bool{
		true,
		value,
	}
}

func (b Bool) IsSet() bool {
	return b.isSet
}

func (b Bool) Value() bool {
	return b.value
}
