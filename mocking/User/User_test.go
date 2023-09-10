package User

import (
	"go.uber.org/mock/gomock"
	"golab.com/m/v2/mocking/mocks"
	"testing"
)
	User
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockUser := mocks.NewMockIUserInterface(mockCtrl)
	testUser = &User.User{IUser: mockUser}

	mockUser.EXPECT().AddUser(1, "sample").Return(nil).Times(1)

	testUser.Use()
}
