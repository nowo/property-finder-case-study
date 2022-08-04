package service

import (
	"property-finder-go-bootcamp-homework/src/config/messages"
	domain "property-finder-go-bootcamp-homework/src/domain/user"
	"property-finder-go-bootcamp-homework/src/domain/user/entity"
	"property-finder-go-bootcamp-homework/src/domain/user/repository"
	"property-finder-go-bootcamp-homework/src/dto/general"
	_jwt "property-finder-go-bootcamp-homework/src/pkg/jwt"
)

type UserService struct {
	Repo repository.IRepository
	jwt  _jwt.JWT
}

func New() IUserService {
	return &UserService{
		Repo: repository.New(),
		jwt:  *_jwt.New(),
	}
}

func (u *UserService) Register(_user domain.User) (general.Token, error) {
	emailExist, _ := u.Repo.CheckEmailExists(_user.UserInfo.Email)

	if emailExist {
		return general.Token{}, messages.EMAIL_ALREADY_EXIST
	}

	user := entity.NewUserInfo(_user.GetUserInfo().Firstname, _user.GetUserInfo().Lastname, _user.GetUserInfo().Email, _user.GetUserInfo().Password)
	createResponse, userCreateErr := u.Repo.Create(domain.User{UserInfo: *user})

	if userCreateErr != nil {
		return general.Token{}, userCreateErr
	}
	token := u.jwt.SetUserId(createResponse.ID).CreateToken().GetToken()

	return general.Token{
		Token: token,
	}, nil
}
