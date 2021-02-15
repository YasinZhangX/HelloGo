package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func Handler(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(content)

	return nil
}
