package repositories

type UserRepository interface {
	Add(name string) bool
	Update(id string,name string) bool
}

func NewUserRepository() UserRepository  {
	return &userRepository{}
}
type userRepository struct {}

func (u userRepository) Add(name string) bool {
	return false
}

func (u userRepository) Update(id string,name string) bool {
	return false
}