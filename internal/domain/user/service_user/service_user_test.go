package service_user

import (
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
	"property-finder-go-bootcamp-homework/test/testdata"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	testdata.MockedDB(testdata.CREATE)
	defer testdata.MockedDB(testdata.DROP)

	db := postgres.ConnectDB()

	defer postgres.Disconnect(db)

	testCases := map[string]struct {
		user        user.User
		expected    general.Token
		expectedErr error
	}{
		"success": {
			user: user.User{
				UserInfo: entity_user.UserInfo{
					Firstname: "John",
					Lastname:  "Doe",
					Email:     "erdal@gmail.com",
					Password:  "123456",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.user.UserInfo.GetUserInfoEmail(), func(t *testing.T) {
			service := New()
			token, err := service.Register(tt.user)
			if err != tt.expectedErr {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
			if token != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, token)
			}
		})
	}
}
