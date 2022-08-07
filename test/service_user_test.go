package test

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
	"property-finder-go-bootcamp-homework/internal/domain/user/service_user"
	"property-finder-go-bootcamp-homework/pkg/errors"
	"property-finder-go-bootcamp-homework/test/mocks"
	"testing"
)

func Test_Login_Failed(t *testing.T) {
	Convey("Register Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := service_user.New(mockUserRepository)
		mockUserRepository.EXPECT().GetUserInfoByEmail(mocks.ValidRequestBody.Email).Return(user.User{}, nil)

		Convey("User Cant Go", func() {
			token, err := userService.Login(auth.LoginRequest{
				Email:    mocks.ValidRequestBody.Email,
				Password: mocks.ValidRequestBody.Password,
			})
			So(err, ShouldResemble, messages.INVALID_PASSWORD)
			So(token, ShouldResemble, general.Token{})
		})
	})
}

func Test_Login_Database_Error(t *testing.T) {
	Convey("Login Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := service_user.New(mockUserRepository)
		mockUserRepository.EXPECT().GetUserInfoByEmail(mocks.ValidRequestBody.Email).Return(user.User{}, messages.DATABASE_OPERATION_FAILED)
		Convey("User Can Go", func() {
			token, err := userService.Login(auth.LoginRequest{
				Email:    mocks.ValidRequestBody.Email,
				Password: mocks.ValidRequestBody.Password,
			})
			So(err, ShouldNotBeNil)
			So(token, ShouldNotBeNil)
		})
	})
}

func Test_Login_Success(t *testing.T) {
	Convey("Login Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := service_user.New(mockUserRepository)
		mockUserRepository.EXPECT().GetUserInfoByEmail(mocks.ValidRequestBody.Email).Return(user.User{
			UserInfo: *entity_user.NewUserInfo(mocks.ValidRequestBody.Firstname, mocks.ValidRequestBody.Lastname, mocks.ValidRequestBody.Email, mocks.ValidRequestBody.Password),
		}, nil)
		Convey("User Can Go", func() {
			token, err := userService.Login(auth.LoginRequest{
				Email:    mocks.ValidRequestBody.Email,
				Password: mocks.ValidRequestBody.Password,
			})
			So(err, ShouldBeNil)
			So(token, ShouldNotBeNil)
		})
	})
}

//Todo: register success passwordde basarili oluyor
//Todo: create fonksiyonunda passwordler uymuyor diye patliyor
//func Test_Register_Success(t *testing.T) {
//	Convey("Register Test Integration", t, func() {
//		mockCtrl := gomock.NewController(t)
//		defer mockCtrl.Finish()
//		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
//		userService := service_user.New(mockUserRepository)
//		mockUserRepository.EXPECT().CheckEmailExists(mocks.ValidRequestBody.Email).Return(false, nil)
//		validUser := user.User{
//			UserInfo: *entity_user.NewUserInfo(mocks.ValidRequestBody.Firstname, mocks.ValidRequestBody.Lastname, mocks.ValidRequestBody.Email, mocks.ValidRequestBody.Password),
//		}
//		mockUserRepository.EXPECT().Create(&validUser).Return(validUser, nil)
//		Convey("User Can Go", func() {
//			token, err := userService.Register(user.User{
//				UserInfo: *entity_user.NewUserInfo(mocks.ValidRequestBody.Firstname, mocks.ValidRequestBody.Lastname, mocks.ValidRequestBody.Email, mocks.ValidRequestBody.Password),
//			})
//			So(err, ShouldBeNil)
//			So(token, ShouldNotBeNil)
//		})
//	})
//}

func Test_Register_EmailAlreadyExist(t *testing.T) {
	Convey("Register Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
		userService := service_user.New(mockUserRepository)
		mockUserRepository.EXPECT().CheckEmailExists(mocks.ValidRequestBody.Email).Return(true, nil)
		Convey("User Can Go", func() {
			token, err := userService.Register(user.User{
				UserInfo: *entity_user.NewUserInfo(mocks.ValidRequestBody.Firstname, mocks.ValidRequestBody.Lastname, mocks.ValidRequestBody.Email, mocks.ValidRequestBody.Password),
			})
			So(err, ShouldResemble, errors.NewEmailAlreadyExist(mocks.ValidRequestBody.Email))
			So(token, ShouldResemble, general.Token{})
		})
	})
}

//func Test_Register_Database_Error(t *testing.T) {
//	Convey("Register Test Integration", t, func() {
//		mockCtrl := gomock.NewController(t)
//		defer mockCtrl.Finish()
//		mockUserRepository := mocks.NewMockIRepository(mockCtrl)
//		userService := service_user.New(mockUserRepository)
//		validUser := user.User{
//			UserInfo: *entity_user.NewUserInfo(mocks.ValidRequestBody.Firstname, mocks.ValidRequestBody.Lastname, mocks.ValidRequestBody.Email, mocks.ValidRequestBody.Password),
//		}
//		mockUserRepository.EXPECT().CheckEmailExists(mocks.ValidRequestBody.Email).Return(false, messages.DATABASE_OPERATION_FAILED)
//		mockUserRepository.EXPECT().Create(&validUser)
//		Convey("User Can Go", func() {
//			token, err := userService.Register(user.User{
//				UserInfo: *entity_user.NewUserInfo(mocks.ValidRequestBody.Firstname, mocks.ValidRequestBody.Lastname, mocks.ValidRequestBody.Email, mocks.ValidRequestBody.Password),
//			})
//			So(err, ShouldNotBeNil)
//			So(token, ShouldNotBeNil)
//		})
//	})
//}
