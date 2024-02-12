package handlers

import (
	"godis/resp"
	"godis/supportedtypes"
)

const PING_CMD = "ping"

func ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: supportedtypes.STRING_TYPE, Str: "PONG"}
	}
	return resp.Value{Typ: supportedtypes.STRING_TYPE, Str: args[0].Bulk}
}
