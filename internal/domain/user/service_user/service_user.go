package service_user

import (
	"fmt"
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
	repository2 "property-finder-go-bootcamp-homework/internal/domain/user/repository_user"
	_jwt "property-finder-go-bootcamp-homework/pkg/jwt"
)

type UserService struct {
	Repo repository2.IRepository
	jwt  _jwt.JWT
}

func New() IUserService {
	return &UserService{
		Repo: repository2.New(),
		jwt:  *_jwt.New(),
	}
}

func (u *UserService) Register(_user user.User) (general.Token, error) {
	emailExist, _ := u.Repo.CheckEmailExists(_user.UserInfo.Email)

	if emailExist {
		fmt.Println("Email already exist!!!")
		return general.Token{}, messages.EMAIL_ALREADY_EXIST
	}
	fmt.Println("Email not exist")

	newUser := entity_user.NewUserInfo(_user.GetUserInfo().Firstname, _user.GetUserInfo().Lastname, _user.GetUserInfo().Email, _user.GetUserInfo().Password)
	createResponse, userCreateErr := u.Repo.Create(user.User{UserInfo: *newUser})

	if userCreateErr != nil {
		return general.Token{}, userCreateErr
	}
	token := u.jwt.SetUserId(createResponse.ID).CreateToken().GetToken()

	return general.Token{
		Token: token,
	}, nil
}

func (u *UserService) Login(_dto auth.LoginRequest) (general.Token, error) {
	user, err := u.Repo.GetUserInfoByEmail(_dto.Email)
	fmt.Println(user)
	if err != nil {
		return general.Token{}, err
	}
	if !user.GetUserInfo().ComparePasswords(_dto.Password) {
		fmt.Println(_dto.Password)
		return general.Token{}, messages.INVALID_PASSWORD
	}

	token := u.jwt.SetUserId(user.ID).CreateToken().GetToken()

	return general.Token{
		Token: token,
	}, nil
}
