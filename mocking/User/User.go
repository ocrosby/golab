package User

import (
	"golab.com/m/v2/mocking/IUser"
)

type user struct {
	IUser IUser.IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1, "sample test")
}
