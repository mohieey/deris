package handlers

import (
	"godis/resp"
	"godis/supportedtypes"
)

const HSET_CMD = "hset"

func hset(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: supportedtypes.ERROR_TYPE, Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return resp.Value{Typ: supportedtypes.STRING_TYPE, Str: "OK"}
}
