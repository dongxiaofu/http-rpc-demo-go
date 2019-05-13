package main

import (
	"errors"
	"net/rpc"
	"net/http"
	"fmt"
	"http-rpc-demo-go/myStruct"
)



type Arith int

func (t *Arith) Multiply(args *myStruct.Args, reply *int) error  {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *myStruct.Args, quo *myStruct.Quotient) error  {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func main()  {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}
