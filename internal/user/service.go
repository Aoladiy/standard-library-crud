package user

type Service struct {
	r *Repo
}

func NewService(r *Repo) *Service {
	return &Service{r: r}
}

func (s Service) GetUsers() ([]User, error) {
	return s.r.getUsers()
}

func (s Service) GetUserById(id int) (user User, err error) {
	return s.r.getUserById(id)
}

func (s Service) CreateUser(user User) (id int, err error) {
	return s.r.createUser(user)
}

func (s Service) UpdateUser(user User) (err error) {
	return s.r.updateUser(user)
}

func (s Service) DeleteUserById(id int) (err error) {
	return s.r.deleteUserById(id)
}
