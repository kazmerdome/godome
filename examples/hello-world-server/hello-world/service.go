package helloworld

import "github.com/kazmerdome/godome/pkg/module/provider/service"

type HelloworldService interface {
	SayHello() string
}

type helloworldService struct {
	service.ServiceConfig
}

func newUserService(c service.ServiceConfig) HelloworldService {
	return &helloworldService{ServiceConfig: c}
}

func (r *helloworldService) SayHello() string {
	r.GetLogger().Info("Helloworld Service has been called.")
	return "Hello world"
}
