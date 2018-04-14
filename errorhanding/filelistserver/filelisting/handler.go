package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFieldlist(writer http.ResponseWriter,
	request *http.Request) error {

	path := request.URL.Path[len("/list/"):]
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
