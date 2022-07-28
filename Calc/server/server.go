package main

import (
	"calc/calcpb"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	calcpb.UnimplementedCalcServer
}

func (*Server) Sum(ctx context.Context, req *calcpb.SumRequest) (resp *calcpb.SumResponse, err error) {
	fmt.Println("Server: Calculator Sum is invoked")
	num1 := req.GetNum1()
	num2 := req.GetNum2()

	result := num1 + num2

	resp = &calcpb.SumResponse{
		Num: result,
	}
	return resp, nil
}

func (*Server) PrimeNum(req *calcpb.PrimeNumberRequest, resp calcpb.Calc_PrimeNumServer) error {
	fmt.Println("Server: Calculator Prime Number is invoked")
	num := req.GetNum()
	var primes []int
	for i := 0; i < int(num); i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	for _, val := range primes {
		result := &calcpb.PrimeNumberResponse{
			Num: int64(val),
		}
		resp.Send(result)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*Server) ComputeAvg(stream calcpb.Calc_ComputeAvgServer) error {
	fmt.Println("Server: Calculator Compute Average is invoked")
	var sum, count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected Error: %v", err)
		}
		sum += req.GetNum()
		count++
	}
	avg := sum / count
	return stream.SendAndClose(&calcpb.ComputeAvgResponse{
		Num: avg,
	})
}

func (*Server) FindMaxNum(stream calcpb.Calc_FindMaxNumServer) error {
	fmt.Println("Server: Calculator Find Maximum is invoked")
	max := math.MinInt
	for {
		req, err := stream.Recv()
		num := int(req.GetNum())
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Unexpected Error: %v", err)
			return err
		}
		if max < num {
			max = num
			err = stream.Send(&calcpb.FindMaxNumResponse{
				Num: int64(max),
			})
			if err != nil {
				log.Fatalf("Unexpected error: %v", err)
				return err
			}
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Server starting...")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServer(s, &Server{})

	err = s.Serve(listen)
	if err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
