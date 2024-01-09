package handlers

import (
	"deris/resp"
	"deris/supportedtypes"
)

const GET_CMD = "get"

func get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: supportedtypes.ERROR_TYPE, Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: supportedtypes.NULL_TYPE}
	}

	return resp.Value{Typ: supportedtypes.BULK_TYPE, Bulk: value}
}
