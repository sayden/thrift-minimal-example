package main

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/sayden/thrift/gen-go/hello"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(
		thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	addr := "localhost:3636"

	socket, err := thrift.NewTSocket(addr)
	if err != nil {
		panic(err)
	}

	transport := transportFactory.GetTransport(socket)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		panic(err)
	}

	client := hello.NewHelloClientFactory(transport, protocolFactory)
	request := hello.NewHelloRequest()
	request.Message = "world!"
	response, err := client.Hello(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
