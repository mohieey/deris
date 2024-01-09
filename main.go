package main

import (
	"deris/aof"
	"deris/handlers"
	"deris/resp"
	"deris/supportedtypes"
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

	aof, err := aof.NewAof("db.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aof.Close()

	aof.Read(func(value resp.Value) {
		command := strings.ToLower(value.Array[0].Bulk)
		args := value.Array[1:]

		handler, ok := handlers.Handlers[command]
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
	writer := resp.NewWriter(conn)

	for {
		respDeserializer := resp.NewRespDeserializer(conn)

		val, err := respDeserializer.Read()
		if err != nil {
			log.Fatal(err)
			return
		}

		if val.Typ != supportedtypes.ARRAY_TYPE {
			writer.Write(resp.Value{Typ: supportedtypes.STRING_TYPE, Str: "ivalid request, expected an array"})
			continue
		}

		if len(val.Array) == 0 {
			writer.Write(resp.Value{Typ: supportedtypes.STRING_TYPE, Str: "ivalid request, expected a non empty array"})
			continue
		}

		command := strings.ToLower(val.Array[0].Bulk)
		args := val.Array[1:]

		handler, ok := handlers.Handlers[command]
		if !ok {
			writer.Write(resp.Value{Typ: supportedtypes.BULK_TYPE, Bulk: fmt.Sprintf("Invalid command: %s", command)})
			continue
		}

		if command == handlers.SET_CMD || command == handlers.HSET_CMD {
			aof.Write(val)
		}

		printVal(&val)

		result := handler(args)
		writer.Write(result)
	}

}

func printVal(v *resp.Value) {
	fmt.Println("================================")
	fmt.Println("type", v.Typ)
	fmt.Println("string", v.Str)
	fmt.Println("number", v.Num)
	fmt.Println("Bulk", v.Bulk)
	fmt.Println("array", v.Array)
	fmt.Println("array len", len(v.Array))
	fmt.Println("================================")
}
