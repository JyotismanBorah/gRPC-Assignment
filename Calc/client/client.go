package main

import (
	"calc/calcpb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client starting...")
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect: %v", err)
	}
	defer cc.Close()
	c := calcpb.NewCalcClient(cc)

	// Sum(c)

	// PrimeNum(c)

	ComputeAvg(c)

	// FindMaxNum(c)
}

func Sum(c calcpb.CalcClient) {
	fmt.Println("Client: Starting Sum")
	req := calcpb.SumRequest{
		Num1: 10,
		Num2: 20,
	}
	resp, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error in calling Sum function: %v", err)
	}
	fmt.Printf("\nRequest Numbers: %v, %v", req.GetNum1(), req.GetNum2())
	fmt.Printf("\nResponse: %v\n", resp.GetNum())
}

func PrimeNum(c calcpb.CalcClient) {
	fmt.Println("Client: Prime Number is called")
	req := calcpb.PrimeNumberRequest{
		Num: 50,
	}
	respStream, err := c.PrimeNum(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error in calling Prime Number: %v", err)
	}
	fmt.Printf("Request Number: %v\n", req.GetNum())

	for {
		resp, err := respStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error in receiving server stream: %v", err)
		}
		log.Printf("Response: %v", resp.GetNum())
	}
}

func ComputeAvg(c calcpb.CalcClient) {
	fmt.Println("Client: Compute Avg is invoked")
	reqStream, err := c.ComputeAvg(context.Background())
	if err != nil {
		log.Fatalf("Error in client side streaming: %v", err)
	}
	requests := []*calcpb.ComputeAvgRequest{
		&calcpb.ComputeAvgRequest{
			Num: 10,
		},
		&calcpb.ComputeAvgRequest{
			Num: 12,
		},
		&calcpb.ComputeAvgRequest{
			Num: 14,
		},
		&calcpb.ComputeAvgRequest{
			Num: 16,
		},
		&calcpb.ComputeAvgRequest{
			Num: 18,
		},
		&calcpb.ComputeAvgRequest{
			Num: 20,
		},
		&calcpb.ComputeAvgRequest{
			Num: 22,
		},
		&calcpb.ComputeAvgRequest{
			Num: 24,
		},
	}
	for _, req := range requests {
		fmt.Printf("Sending Request: Input = %v\n", req.GetNum())
		reqStream.Send(req)
		time.Sleep(1 * time.Second)
	}
	resp, err := reqStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error in receiving response: %v", err)
	}
	log.Printf("Response: %v", resp.GetNum())
}

func FindMaxNum(c calcpb.CalcClient) {
	fmt.Println("Client: Find Max is invoked")
	requests := []*calcpb.FindMaxNumRequest{
		&calcpb.FindMaxNumRequest{
			Num: 1,
		},
		&calcpb.FindMaxNumRequest{
			Num: 3,
		},
		&calcpb.FindMaxNumRequest{
			Num: 7,
		},
		&calcpb.FindMaxNumRequest{
			Num: 9,
		},
		&calcpb.FindMaxNumRequest{
			Num: 2,
		},
		&calcpb.FindMaxNumRequest{
			Num: 5,
		},
		&calcpb.FindMaxNumRequest{
			Num: 22,
		},
		&calcpb.FindMaxNumRequest{
			Num: 15,
		},
		&calcpb.FindMaxNumRequest{
			Num: 21,
		},
		&calcpb.FindMaxNumRequest{
			Num: 19,
		},
		&calcpb.FindMaxNumRequest{
			Num: 108,
		},
	}
	stream, err := c.FindMaxNum(context.Background())
	if err != nil {
		log.Printf("Error is streaming: %v", err)
	}

	waitchan := make(chan struct{})

	go func(requests []*calcpb.FindMaxNumRequest) {
		for _, req := range requests {
			err := stream.Send(req)
			if err != nil {
				log.Printf("Error in sending request stream: %v", err)
			}
			fmt.Printf("\nRequest: %v\t", req.GetNum())
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}(requests)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitchan)
				return
			}
			if err != nil {
				log.Printf("Error in receiving response stream: %v", err)
			}
			fmt.Printf("Response: %v", resp.GetNum())
		}
	}()

	<-waitchan
}
