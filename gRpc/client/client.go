package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"zRpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	single(client)
}

func single(client proto.HelloServiceClient) {
	reply, err := client.Hello(context.Background(), &proto.String{Value: "zs"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}

func stream(client proto.HelloServiceClient) {
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := stream.Send(&proto.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
