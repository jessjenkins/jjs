package service

type Service struct {
}

type SubService struct {
	Name string
}

type Servicer interface {
}

func NewService() Service {
	return Service{}
}

func (s *Service) AddSubService(name string, subservice Servicer, dep ...*SubService) *SubService {
	return &SubService{Name: name}
}

func (s *Service) Run() {

}
