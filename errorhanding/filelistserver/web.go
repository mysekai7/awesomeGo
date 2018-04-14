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
		//handler处理业务逻辑
		err := handler(writer, request)
		//错误处理
		if err != nil {
			log.Warn("Error handling request: %s", err.Error())
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

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFieldlist))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
