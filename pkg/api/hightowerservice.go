package api

// HightowerService interface
type HightowerService interface {
	NewService(s *HightowerService)
	Name() string
	Ports() []string
	Available() bool
	Before()
	After()
	Start()
	Stop()
	Delete()
}

// NilService is the core interface for running and
// interacting with deployables
type NilService struct {
	name  string
	ports []uint
}

// Name ...
func (s *NilService) Name() string {
	return s.name
}
