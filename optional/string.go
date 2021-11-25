package optional

type NullableString struct {
	Value  string
	AsNull bool
}

func String(v string, as_null bool) NullableString {
	return NullableString{
		Value:  v,
		AsNull: as_null,
	}
}
