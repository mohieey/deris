package main

import "sync"

const (
	PING_CMD    = "ping"
	SET_CMD     = "set"
	GET_CMD     = "get"
	HSET_CMD    = "hset"
	HGET_CMD    = "hget"
	HGETALL_CMD = "hgetall"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

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

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: ERROR_TYPE, str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return Value{typ: STRING_TYPE, str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: ERROR_TYPE, str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return Value{typ: NULL_TYPE}
	}

	return Value{typ: BULK_TYPE, bulk: value}
}

func hgetall(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: ERROR_TYPE, str: "ERR wrong number of arguments for 'hgetall' command"}
	}

	hash := args[0].bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash]
	HSETsMu.RUnlock()

	if !ok {
		return Value{typ: NULL_TYPE}
	}

	values := []Value{}
	for k, v := range value {
		values = append(values, Value{typ: BULK_TYPE, bulk: k})
		values = append(values, Value{typ: BULK_TYPE, bulk: v})
	}

	return Value{typ: ARRAY_TYPE, array: values}
}

var Handlers = map[string]func([]Value) Value{
	PING_CMD:    ping,
	SET_CMD:     set,
	GET_CMD:     get,
	HSET_CMD:    hset,
	HGET_CMD:    hget,
	HGETALL_CMD: hgetall,
}
