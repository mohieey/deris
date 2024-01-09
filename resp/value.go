package resp

import (
	"deris/supportedtypes"
	"strconv"
)

type Value struct {
	Typ   string
	Str   string
	Num   int
	Bulk  string
	Array []Value
}

func (v Value) marshalArray() []byte {
	len := len(v.Array)
	var bytes []byte
	bytes = append(bytes, supportedtypes.ARRAY_SYMBOL)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len; i++ {
		bytes = append(bytes, v.Array[i].Marshal()...)
	}

	return bytes
}

func (v Value) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, supportedtypes.BULK_SYMBOL)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, supportedtypes.STRING_SYMBOL)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallError() []byte {
	var bytes []byte
	bytes = append(bytes, supportedtypes.ERROR_SYMBOL)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallNull() []byte {
	return []byte("$-1\r\n")
}

func (v *Value) Marshal() []byte {
	switch v.Typ {
	case "array":
		return v.marshalArray()
	case "Bulk":
		return v.marshalBulk()
	case "string":
		return v.marshalString()
	case "null":
		return v.marshallNull()
	case "error":
		return v.marshallError()
	default:
		return []byte{}
	}
}
