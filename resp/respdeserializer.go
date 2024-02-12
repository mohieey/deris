package resp

import (
	"bufio"
	"fmt"
	"godis/supportedtypes"
	"io"
	"strconv"
)

type RespDeserializer struct {
	reader *bufio.Reader
}

func NewRespDeserializer(reader io.Reader) *RespDeserializer {
	return &RespDeserializer{reader: bufio.NewReader(reader)}
}

func (r *RespDeserializer) readLine() ([]byte, int, error) {
	bytesRead := 0
	var line []byte

	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		bytesRead += 1
		line = append(line, b)

		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], bytesRead, nil
}

func (r *RespDeserializer) readInteger() (int, int, error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	i, err := strconv.Atoi(string(line))
	if err != nil {
		return 0, n, err
	}

	return i, n, nil
}

func (r *RespDeserializer) readArray() (Value, error) {
	v := Value{Typ: supportedtypes.ARRAY_TYPE}

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.Array = make([]Value, 0, len)
	for i := 0; i < len; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}

		v.Array = append(v.Array, val)
	}

	return v, nil
}

func (r *RespDeserializer) readBulk() (Value, error) {
	v := Value{Typ: supportedtypes.BULK_TYPE}

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	Bulk := make([]byte, len)

	r.reader.Read(Bulk)

	v.Bulk = string(Bulk)

	r.readLine()

	return v, nil
}

func (r *RespDeserializer) Read() (Value, error) {
	_type, err := r.reader.ReadByte()

	if err != nil {
		return Value{}, err
	}

	switch _type {
	case supportedtypes.ARRAY_SYMBOL:
		return r.readArray()
	case supportedtypes.BULK_SYMBOL:
		return r.readBulk()
	default:
		fmt.Printf("Unknown type: %v", string(_type))
		return Value{}, nil
	}
}
