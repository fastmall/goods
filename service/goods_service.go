package service

import (
	context "context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/fastmall/goods/api"
)

type GoodsService struct {
	api.UnimplementedGoodsServiceServer
}

func (s GoodsService) CreateGoodsItem(ctx context.Context, req *api.CreateGoodsItemRequest) (*api.CreateGoodsItemResponse, error) {
	logger.Infof("CreateGoodsItem %#v", req)

	res := &api.CreateGoodsItemResponse{
		CustomerId: 0,
		Token:      "",
	}
	return res, nil
}
