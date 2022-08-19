package jjs

type Runner struct {
}

type Service struct {
	Name string
}

type Servicer interface {
}

func NewRunner() Runner {
	return Runner{}
}

func (s *Runner) AddService(name string, subservice Servicer, dep ...*Service) *Service {
	return &Service{Name: name}
}

func (s *Runner) Run() {

}
