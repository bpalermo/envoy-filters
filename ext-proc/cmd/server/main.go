package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type ExternalProcessorServer struct {
	pb.UnimplementedExternalProcessorServer
}

func (e *ExternalProcessorServer) Process(stream pb.ExternalProcessor_ProcessServer) (err error) {
	for {
		ctx := stream.Context()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive stream request: %v", err)
		}

		resp := &pb.ProcessingResponse{}
		switch value := req.Request.(type) {
		case *pb.ProcessingRequest_RequestHeaders:
			resp = &pb.ProcessingResponse{
				Response: &pb.ProcessingResponse_RequestHeaders{},
			}
			break
		case *pb.ProcessingRequest_RequestBody:
			resp = &pb.ProcessingResponse{
				Response: &pb.ProcessingResponse_RequestBody{},
			}
			break
		case *pb.ProcessingRequest_RequestTrailers:
			resp = &pb.ProcessingResponse{
				Response: &pb.ProcessingResponse_RequestTrailers{},
			}
			break
		case *pb.ProcessingRequest_ResponseHeaders:
			resp = &pb.ProcessingResponse{
				Response: &pb.ProcessingResponse_ResponseHeaders{},
			}
			break
		case *pb.ProcessingRequest_ResponseBody:
			resp = &pb.ProcessingResponse{
				Response: &pb.ProcessingResponse_ResponseBody{},
			}
			break
		case *pb.ProcessingRequest_ResponseTrailers:
			resp = &pb.ProcessingResponse{
				Response: &pb.ProcessingResponse_ResponseTrailers{},
			}
			break
		default:
			log.Printf("Unknown Request type %v\n", value)
		}
		if err := stream.Send(resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalCtx, signalCtxStop := signal.NotifyContext(ctx,
		os.Interrupt,    // interrupt = SIGINT = Ctrl+C
		syscall.SIGQUIT, // Ctrl-\
		syscall.SIGTERM, // "the normal way to politely ask a program to terminate"
		syscall.SIGKILL, //
	)
	defer signalCtxStop()

	go func() {
		<-signalCtx.Done()
		grpcServer.GracefulStop()
	}()

	pb.RegisterExternalProcessorServer(grpcServer, &ExternalProcessorServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
