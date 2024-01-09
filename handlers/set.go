package handlers

import (
	"deris/resp"
	"deris/supportedtypes"
)

const SET_CMD = "set"

func set(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: supportedtypes.ERROR_TYPE, Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return resp.Value{Typ: supportedtypes.STRING_TYPE, Str: "OK"}
}
