package mock

import "fmt"

type Retriver struct {
	Contents string
}

func (r *Retriver) String() string {
	return fmt.Sprintf("Retriver: {Contents=%s}", r.Contents)
}

func (r *Retriver) Get(url string) string {
	return r.Contents
}

func (r *Retriver) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
