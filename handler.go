package main

const (
	PING_CMD = "ping"
)

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: STRING_TYPE, str: "PONG"}
	}
	return Value{typ: STRING_TYPE, str: args[0].bulk}
}

var Handlers = map[string]func([]Value) Value{
	PING_CMD: ping,
}
