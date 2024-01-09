package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const PORT = ":6379"

func main() {
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("deris is listening on:", PORT)

	aof, err := NewAof("db.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aof.Close()

	aof.Read(func(value Value) {
		command := strings.ToLower(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command:", command)
			return
		}

		handler(args)
	})

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	writer := NewWriter(conn)

	for {
		resp := NewResp(conn)

		val, err := resp.Read()
		if err != nil {
			log.Fatal(err)
			return
		}

		if val.typ != ARRAY_TYPE {
			writer.Write(Value{typ: STRING_TYPE, str: "ivalid request, expected an array"})
			continue
		}

		if len(val.array) == 0 {
			writer.Write(Value{typ: STRING_TYPE, str: "ivalid request, expected a non empty array"})
			continue
		}

		command := strings.ToLower(val.array[0].bulk)
		args := val.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			writer.Write(Value{typ: BULK_TYPE, bulk: fmt.Sprintf("Invalid command: %s", command)})
			continue
		}

		if command == SET_CMD || command == HSET_CMD {
			aof.Write(val)
		}

		printVal(&val)

		result := handler(args)
		writer.Write(result)
	}

}

func printVal(v *Value) {
	fmt.Println("================================")
	fmt.Println("type", v.typ)
	fmt.Println("string", v.str)
	fmt.Println("number", v.num)
	fmt.Println("bulk", v.bulk)
	fmt.Println("array", v.array)
	fmt.Println("array len", len(v.array))
	fmt.Println("================================")
}
