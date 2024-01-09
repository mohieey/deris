package main

import (
	"fmt"
	"log"
	"net"
)

const PORT = ":6379"

func main() {
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("deris is listening on:", PORT)

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		resp := NewResp(conn)

		val, err := resp.Read()
		if err != nil {
			log.Fatal(err)
			return
		}

		printVal(&val)

		conn.Write([]byte("+OK\r\n"))
	}

	// input := "$5\r\nAhmed\r\n"

	// reader := bufio.NewReader(strings.NewReader(input))
	// b, _ := reader.ReadByte()

	// if b != '$' {
	// 	fmt.Println("Invalid type, expecting bulk strings only")
	// 	os.Exit(1)
	// }

	// size, _ := reader.ReadByte()

	// strSize, _ := strconv.Atoi(string(size))

	// // consume /r/n
	// reader.ReadByte()
	// reader.ReadByte()

	// name := make([]byte, strSize)
	// reader.Read(name)

	// fmt.Println(string(name))

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
