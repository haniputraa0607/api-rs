package userservice

import (
	userrepository "api-rs/repositories/user"
	"api-rs/schemas"
	"api-rs/utility"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	ListUser() (response []*schemas.ListUserResponse, err error)
	Login(username string, password string) (*schemas.LoginResponse, error)
}

type userService struct {
	userRepository userrepository.UserRepository
}

func NewUserService(
	userRepository userrepository.UserRepository,
) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) ListUser() (response []*schemas.ListUserResponse, err error) {
	users, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		response = append(response, &schemas.ListUserResponse{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	return response, nil
}

func (s *userService) Login(username string, password string) (*schemas.LoginResponse, error) {
	user, err := s.userRepository.GetUser(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := utility.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	response := schemas.LoginResponse{
		Token: *token,
	}

	return &response, nil
}
