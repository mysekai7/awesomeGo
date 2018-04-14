package main

import (
	"awesomeProject1/errorhanding/filelistserver/filelisting"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		//	防止崩溃
		defer func() {
			if r := recover(); r != nil {
				log.Warn("%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		//handler处理业务逻辑
		err := handler(writer, request)
		//错误处理
		if err != nil {
			log.Warn("Error handling request: %s", err.Error())

			if userErr, ok := err.(userError); ok { //.() Type Assertion判断类型
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest) //显示在header的状态码
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//自定义错误
type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFieldlist))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
