package handlers

import (
	"godis/resp"
	"godis/supportedtypes"
)

const HGET_CMD = "hget"

func hget(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: supportedtypes.ERROR_TYPE, Str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: supportedtypes.NULL_TYPE}
	}

	return resp.Value{Typ: supportedtypes.BULK_TYPE, Bulk: value}
}
