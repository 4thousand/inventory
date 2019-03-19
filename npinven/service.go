package NPINven

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	GenDocNoInven(Type string, Search string, Branch string) (interface{}, error)
}

func (s *service) GenDocNoInven(Type string, Search string, Branch string) (interface{}, error) {
	repo, err := s.repo.GenDocNoInven(Type, Search, Branch)
	if err != nil {
		return nil, err
	}
	return repo, nil
}
