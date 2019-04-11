package dog

import "awesomeGo/manager"

func NewDog() *Dog {
	return &Dog{}
}

type Dog struct {
}

func (d *Dog) Say() {
	manager.SharedInstance().Verify("dog")
}
