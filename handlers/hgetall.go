package handlers

import (
	"godis/resp"
	"godis/supportedtypes"
)

const HGETALL_CMD = "hgetall"

func hgetall(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: supportedtypes.ERROR_TYPE, Str: "ERR wrong number of arguments for 'hgetall' command"}
	}

	hash := args[0].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash]
	HSETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: supportedtypes.NULL_TYPE}
	}

	values := []resp.Value{}
	for k, v := range value {
		values = append(values, resp.Value{Typ: supportedtypes.BULK_TYPE, Bulk: k})
		values = append(values, resp.Value{Typ: supportedtypes.BULK_TYPE, Bulk: v})
	}

	return resp.Value{Typ: supportedtypes.ARRAY_TYPE, Array: values}
}
