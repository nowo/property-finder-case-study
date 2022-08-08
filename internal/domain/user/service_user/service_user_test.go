package service_user

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
	"property-finder-go-bootcamp-homework/pkg/errors"
	"property-finder-go-bootcamp-homework/test_data"
	"property-finder-go-bootcamp-homework/test_data/mocks"
	"testing"
)

func TestLoginFailedWithInvalidPassword(t *testing.T) {
	Convey("Given that i tried to login with invalid user", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := New(mockUserRepository)
		mockUserRepository.EXPECT().GetUserInfoByEmail(test_data.ValidRequestBody.Email).Return(user.User{}, nil)

		Convey("Then i get invalid password error", func() {
			request := auth.LoginRequest{
				Email:    test_data.ValidRequestBody.Email,
				Password: test_data.ValidRequestBody.Password,
			}
			token, err := userService.Login(request)
			So(err, ShouldResemble, messages.INVALID_PASSWORD)
			So(token, ShouldResemble, general.Token{})
		})
	})
}

func Test_LoginDatabaseError(t *testing.T) {
	Convey("Given that when i tried to login, database bugged", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := New(mockUserRepository)
		mockUserRepository.EXPECT().GetUserInfoByEmail(test_data.ValidRequestBody.Email).Return(user.User{}, messages.DATABASE_OPERATION_FAILED)
		Convey("Then i get database error", func() {
			token, err := userService.Login(auth.LoginRequest{
				Email:    test_data.ValidRequestBody.Email,
				Password: test_data.ValidRequestBody.Password,
			})
			So(err, ShouldNotBeNil)
			So(token, ShouldNotBeNil)
		})
	})
}

func Test_LoginSuccess(t *testing.T) {
	Convey("Given that i try to login with valid user", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := New(mockUserRepository)
		mockUserRepository.EXPECT().GetUserInfoByEmail(test_data.ValidRequestBody.Email).Return(user.User{
			UserInfo: *entity_user.NewUserInfo(test_data.ValidRequestBody.Firstname, test_data.ValidRequestBody.Lastname, test_data.ValidRequestBody.Email, test_data.ValidRequestBody.Password),
		}, nil)
		Convey("Then i logged in successfully", func() {
			token, err := userService.Login(auth.LoginRequest{
				Email:    test_data.ValidRequestBody.Email,
				Password: test_data.ValidRequestBody.Password,
			})
			So(err, ShouldBeNil)
			So(token, ShouldNotBeNil)
		})
	})
}

func Test_RegisterEmailExist(t *testing.T) {
	Convey("Given that i tried to register with already registered email", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := New(mockUserRepository)
		mockUserRepository.EXPECT().CheckEmailExists(test_data.ValidRequestBody.Email).Return(true, nil)
		Convey("Then i get email already exist error", func() {
			token, err := userService.Register(user.User{
				UserInfo: *entity_user.NewUserInfo(test_data.ValidRequestBody.Firstname, test_data.ValidRequestBody.Lastname, test_data.ValidRequestBody.Email, test_data.ValidRequestBody.Password),
			})
			So(err, ShouldResemble, errors.NewEmailAlreadyExist(test_data.ValidRequestBody.Email))
			So(token, ShouldResemble, general.Token{})
		})
	})
}
