package optional

type NullableInt struct {
	Value  int64
	AsNull bool
}

func Int(v int64, as_null bool) NullableInt {
	return NullableInt{
		Value:  v,
		AsNull: as_null,
	}
}
