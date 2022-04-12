package main

import (
	"context"
	"io"
	"log"

	pb "github.com/joelgarciajr84/go-grpc-stream-client/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// So call me maybe :)

	conn, err := grpc.Dial(":50005", grpc.WithTransportCredentials(insecure.NewCredentials())) // Don't try this shit at home ...
	if err != nil {
		log.Fatalf("Error connecting  %v", err)
	}

	// Follow the flow
	client := pb.NewStreamServiceClient(conn)
	in := &pb.Request{Id: 1}
	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true // Send true to channel, means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("Error receiveing %v", err)
			}
			log.Printf("Response received: %s", resp.Result)
		}
	}()

	<-done //Waiting ...
	log.Printf("finished")
}
