package mock

type Retriver struct {
	Contents string
}

func (r *Retriver) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
func (r *Retriver) Get(url string) string {
	return r.Contents
}
