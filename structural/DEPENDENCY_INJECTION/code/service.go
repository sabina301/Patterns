package code

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetUserById(id int) string {
	return s.rep.GetUser(id)
}
