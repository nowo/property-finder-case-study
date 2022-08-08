package service_user

import (
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
	"property-finder-go-bootcamp-homework/internal/domain/user/repository_user"

	"property-finder-go-bootcamp-homework/pkg/errors"
	_jwt "property-finder-go-bootcamp-homework/pkg/jwt"
)

//UserService struct contains all methods that are required to implement by service_user.
type UserService struct {
	Repo repository_user.IUserRepository
	jwt  _jwt.JWT
}

//NewUserService creates a new UserService
func New(repo repository_user.IUserRepository) IUserService {
	return &UserService{
		Repo: repo,
		jwt:  *_jwt.New(),
	}
}

//Register creates a new user by checking email is exist
func (u *UserService) Register(_user user.User) (general.Token, error) {
	emailExist, _ := u.Repo.CheckEmailExists(_user.UserInfo.Email)

	if emailExist {
		return general.Token{}, errors.NewEmailAlreadyExist(_user.UserInfo.Email)
	}

	newUser := entity_user.NewUserInfo(_user.UserInfo.Firstname, _user.UserInfo.Lastname, _user.UserInfo.Email, _user.UserInfo.Password)
	createResponse, err := u.Repo.Create(user.User{UserInfo: *newUser})

	if err != nil {
		return general.Token{}, err
	}
	token := u.jwt.SetUserId(createResponse.ID).CreateToken().GetToken()

	return general.Token{
		Token: token,
	}, nil
}

//Login checks if user exists and returns token
func (u *UserService) Login(_dto auth.LoginRequest) (general.Token, error) {
	user, err := u.Repo.GetUserInfoByEmail(_dto.Email)
	if err != nil {
		return general.Token{}, err
	}
	if !user.UserInfo.ComparePasswords(_dto.Password) {
		return general.Token{}, messages.INVALID_PASSWORD
	}

	token := u.jwt.SetUserId(user.ID).CreateToken().GetToken()

	return general.Token{
		Token: token,
	}, nil
}
