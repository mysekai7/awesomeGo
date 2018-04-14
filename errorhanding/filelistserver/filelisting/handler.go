package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFieldlist(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("Path mast start" +
			"with" + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer,
		//	err.Error(),
		//	http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
