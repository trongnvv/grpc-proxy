package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
	pb "grpc-proxy/proto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedTestServiceServer
}

func (s *server) CallTest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Response{Number: "01234"}, nil
}

type Handler struct {
}

func (h *Handler) TagRPC(context.Context, *stats.RPCTagInfo) context.Context {
	log.Println("TagRPC")
	return context.Background()
}

// HandleRPC processes the RPC stats.
func (h *Handler) HandleRPC(context.Context, stats.RPCStats) {

	//log.Println("HandleRPC")
}

func (h *Handler) TagConn(context.Context, *stats.ConnTagInfo) context.Context {

	log.Println("connect")
	return context.Background()
}

// HandleConn processes the Conn stats.
func (h *Handler) HandleConn(c context.Context, s stats.ConnStats) {
	switch s.(type) {
	case *stats.ConnEnd:
		log.Println("end")
		//fmt.Printf("client %d disconnected", s.userIdMap[ctx.Value("user_counter")])
		break
	}
}

// init your grpc server,like this:
func main() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.StatsHandler(&Handler{}))

	pb.RegisterTestServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
