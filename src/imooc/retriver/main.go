package main

import (
	"fmt"
	"imooc/retriver/mock"
	real2 "imooc/retriver/real"
	"time"
)

type Retriver interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

type RetriverPoster interface {
	Retriver
	Poster
}

const url = "http://www.imooc.com"

func download(r Retriver) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}
func main() {
	var r Retriver
	retriver := mock.Retriver{"This is a fake imooc.com"}
	//fmt.Printf("%T %v\n",r,r)
	r = &retriver
	inspect(r)
	r = &real2.Retriver{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))

	//TYPE ASSERTION
	//realRetriver := r.(*real2.Retriver)
	//retriver := realRetriver
	//fmt.Println(retriver.TimeOut)
	if mockRetriver, ok := r.(*mock.Retriver); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not a mock retriver")
	}
	fmt.Println("Try a session")
	fmt.Println(session(&retriver))

}

func inspect(r Retriver) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriver:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
