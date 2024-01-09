package handlers

import "deris/resp"

var Handlers = map[string]func([]resp.Value) resp.Value{
	PING_CMD:    ping,
	SET_CMD:     set,
	GET_CMD:     get,
	HSET_CMD:    hset,
	HGET_CMD:    hget,
	HGETALL_CMD: hgetall,
}
