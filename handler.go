package main

import "sync"

const (
	PING_CMD = "ping"
	SET_CMD  = "set"
	GET_CMD  = "get"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: STRING_TYPE, str: "PONG"}
	}
	return Value{typ: STRING_TYPE, str: args[0].bulk}
}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: ERROR_TYPE, str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].bulk
	value := args[1].bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{typ: STRING_TYPE, str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: ERROR_TYPE, str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return Value{typ: NULL_TYPE}
	}

	return Value{typ: BULK_TYPE, bulk: value}
}

var Handlers = map[string]func([]Value) Value{
	PING_CMD: ping,
	SET_CMD:  set,
	GET_CMD:  get,
}
