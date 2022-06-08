package delete

type Service interface {
	DeleteService(inputDelete *InputDelete) (uint, string)
}

type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteService(inputDelete *InputDelete) (uint, string) {
	return s.repository.DeleteRepository(inputDelete)
}
