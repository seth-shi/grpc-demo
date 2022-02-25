package main

import (
	"context"
	"github.com/seth-shi/grpc-demo/enums"
	"github.com/seth-shi/grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type server struct {
	pb.UnimplementedOrderServer

	// 为了简单, 当住数据库
	data map[int64][]*pb.OrderData
	// 订单的自增 ID
	id int64
	sync.Mutex
}

func newOrderServer() *server {
	return &server{data: make(map[int64][]*pb.OrderData)}
}

func (s *server) Index(c context.Context, r *pb.OrderIndexRequest) (*pb.OrderIndexResponse, error) {

	return &pb.OrderIndexResponse{Data: s.data[r.UserId]}, nil
}

func (s *server) Store(c context.Context, r *pb.OrderStoreRequest) (*pb.OrderData, error) {
	s.Lock()
	defer s.Unlock()

	s.id++
	d := &pb.OrderData{
		No:        s.id,
		Amount:    r.Amount,
		Number:    r.GoodsNumber,
		GoodsId:   r.GoodsId,
		GoodsName: r.GoodsName,
	}

	s.data[r.UserId] = append(s.data[r.UserId], d)

	return d, nil
}

func main() {

	lis, err := net.Listen("tcp", enums.OrderHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServer(s, newOrderServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
