package main

import (
	"os"
	"fmt"
	"net/rpc"
	"log"
	"http-rpc-demo-go/myStruct"
	"strconv"
)

func main()  {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ", os.Args[0], "server 3 4")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	a, _ := strconv.Atoi(os.Args[2])
	b, _ := strconv.Atoi(os.Args[3])
	args := myStruct.Args{a, b}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot myStruct.Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
