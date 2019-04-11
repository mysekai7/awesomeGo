package main

import (
	"awesomeGo/manager"
	"awesomeGo/user"
	"fmt"
	"github.com/pkg/errors"
)

func main() {

	u := user.User{"xiaobai"}
	u.Verify()

	manager.SharedInstance().Save("hello world")

	u.Verify()

	fmt.Println(errors.Wrap(errors.New("err"), "err message"))

}
