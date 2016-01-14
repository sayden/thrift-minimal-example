/*
 * Mario Castro 2016
 * Server
 */
package main

import (
	"fmt"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/sayden/thrift/gen-go/hello"
)

// HelloHandler is the struct that implement the public endpoint
type HelloHandler struct {
}

//Hello is a public method implementing the
func (h *HelloHandler) Hello(r *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := hello.NewHelloResponse()
	resp.Message = fmt.Sprintf("hello %s", r.Message)
	fmt.Printf("%d - %s\n", time.Now().UnixNano(), resp.Message)
	return resp, nil
}

func main() {
	transport, err := thrift.NewTServerSocket("localhost:3636")
	if err != nil {
		panic(err)
	}

	proc := hello.NewHelloProcessor(&HelloHandler{})
	server := thrift.NewTSimpleServer4(
		proc,
		transport,
		thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory()),
		thrift.NewTBinaryProtocolFactoryDefault())

	println(server.Serve())
}
