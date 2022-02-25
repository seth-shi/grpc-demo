package main

import (
	"context"
	"errors"
	"github.com/seth-shi/grpc-demo/enums"
	"github.com/seth-shi/grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGoodsServer

	// 为了简单, 当住数据库
	data map[int64]*pb.GoodsData
}

func newGoodsServer() *server {
	return &server{data: map[int64]*pb.GoodsData{
		1: {
			Id:     1,
			Name:   "橘子",
			Amount: 9.9,
		},
		2: {
			Id:     2,
			Name:   "香蕉",
			Amount: 8.8,
		},
	}}
}

func (s *server) Show(ctx context.Context, in *pb.GoodsShowRequest) (*pb.GoodsData, error) {

	u, e := s.data[in.Id]
	if !e {
		return nil, errors.New("无此商品")
	}

	return u, nil
}

func main() {

	lis, err := net.Listen("tcp", enums.GoodsHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGoodsServer(s, newGoodsServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
