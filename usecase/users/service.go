package users

type Service struct {
	Repository
	JWTSecret string
}

func LoadService(repository Repository, jwtSecret string) *Service {
	return &Service{
		Repository: repository,
		JWTSecret:  jwtSecret,
	}
}
