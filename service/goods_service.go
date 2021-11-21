package service

import (
	context "context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/fastmall/goods/api"
)

type GoodsService struct {
	api.UnimplementedGoodsServiceServer
}

func (s GoodsService) CreateItem(ctx context.Context, req *api.CreateItemRequest) (*api.CreateItemResponse, error) {
	logger.Infof("CreateGoodsItem %#v", req)

	res := &api.CreateItemResponse{
		CustomerId: 0,
		Token:      "",
	}
	return res, nil
}

func (s GoodsService) GetItemDetail(ctx context.Context, req *api.GetItemDetailRequest) (*api.GetItemDetailResponse, error) {
	res := api.GetItemDetailResponse{
		ItemId: req.ItemId,
		Title:  "goods title",
		Desc:   "goods describe",
	}
	return &res, nil
}
