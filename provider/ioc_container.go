package provider

import (
	"github.com/vu-nhan/gin-dig-grpc/handlers"
	"github.com/vu-nhan/gin-dig-grpc/repositories"
	"github.com/vu-nhan/gin-dig-grpc/services"
	"go.uber.org/dig"
)

type IocContainer struct {
	DigContainer *dig.Container
}

func NewIocContainer() *IocContainer {
	return &IocContainer{
		DigContainer: dig.New(),
	}
}

func (c *IocContainer) InitializeDependency() {
	_ = c.DigContainer.Provide(repositories.NewVehicleRepository)
	_ = c.DigContainer.Provide(repositories.NewCustomerRepository)

	_ = c.DigContainer.Provide(services.NewVehicleService)
	_ = c.DigContainer.Provide(services.NewCustomerService)

	_ = c.DigContainer.Provide(handlers.NewVehicleHandler)
	_ = c.DigContainer.Provide(handlers.NewCustomerHandler)
}

func (c *IocContainer) Invoke(injectedInterface interface{}) error {
	return c.DigContainer.Invoke(injectedInterface)
}
