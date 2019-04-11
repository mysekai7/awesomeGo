package user

import (
	dog2 "awesomeGo/dog"
	"awesomeGo/manager"
)

type User struct {
	Name string
}

func (u *User) Verify() {
	manager.SharedInstance().Verify(u.Name)

	dog := dog2.NewDog()
	dog.Say()
}
