package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seth-shi/grpc-demo/enums"
	"github.com/seth-shi/grpc-demo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

var goodsClient pb.GoodsClient
var orderClient pb.OrderClient

func main() {

	goodsClient = createGoodsClient()
	orderClient = createOrderClient()

	// 注册 HTTP 路由
	r := gin.Default()
	r.GET("/orders", orderIndex)
	r.POST("/orders", createOrder)
	r.Run(enums.HttpHost)
}

func createGoodsClient() pb.GoodsClient {
	conn, err := grpc.Dial(enums.GoodsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewGoodsClient(conn)
}

func createOrderClient() pb.OrderClient {
	conn, err := grpc.Dial(enums.OrderHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewOrderClient(conn)
}

type createOrderRequest struct {
	UserId  int64 `json:"user_id" form:"user_id" binding:"required"`
	GoodsId int64 `json:"goods_id" form:"goods_id" binding:"required"`
	Number  int64 `json:"number" form:"number" binding:"required"`
}

type createOrderResponse struct {
	No      int64   `json:"no"`
	Name    string  `json:"name"`
	Number  int64   `json:"number"`
	GoodsId int64   `json:"goods_id"`
	Amount  float64 `json:"amount"`
}

func createOrder(c *gin.Context) {

	// 获取输入的参数
	var req createOrderRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	goods, err := goodsClient.Show(c.Request.Context(), &pb.GoodsShowRequest{Id: req.GoodsId})
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	res, err := orderClient.Store(c.Request.Context(), &pb.OrderStoreRequest{
		GoodsId:     goods.Id,
		UserId:      req.UserId,
		GoodsNumber: req.Number,
		GoodsName:   goods.Name,
		Amount:      float64(req.Number) * goods.Amount,
	})
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": createOrderResponse{
			No:      res.No,
			Name:    res.GoodsName,
			Number:  res.Number,
			Amount:  res.Amount,
			GoodsId: res.GoodsId,
		},
	})
}

func orderIndex(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "error": "无效的用户"})
		return
	}

	res, err := orderClient.Index(c.Request.Context(), &pb.OrderIndexRequest{UserId: userId})
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": res.Data,
	})
}
