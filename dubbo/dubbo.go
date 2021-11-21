package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	customerApi "github.com/fastmall/customer/api"
	"github.com/fastmall/goods/service"
)

var CustomerService = &customerApi.CustomerServiceClientImpl{}

func StartDubbo() {
	config.SetConsumerService(CustomerService)
	config.SetProviderService(&service.GoodsService{})
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
