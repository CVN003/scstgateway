package main

import (
	"fmt"
	"net"

	"github.com/CVN003/scstgateway/scst"
	"google.golang.org/grpc"
)

func main() {
	var c chan int
	go func() {
		lis, err := net.Listen("tcp", ":55101")
		if err != nil {
			panic(err)
		}
		ser := grpc.NewServer()
		scst.RegisterSCSTGatewayServer(ser, &scst.Gateway{})
		if err := ser.Serve(lis); err != nil {
			panic(err)
		}
		c <- 1
	}()

	fmt.Printf("scst gateway server start at :55101")
	<-c
}
