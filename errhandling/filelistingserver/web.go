package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"

	"yasinzhang.top/hellogo/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	logger, _ := zap.NewProduction()

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			logger.Sugar().Warnf("Error Handling request: %s", err.Error())

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.Handler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
