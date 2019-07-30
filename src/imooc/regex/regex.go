package main

import (
	"fmt"
	"net/http"
	"regexp"
)

//const text = "my email is allanyang@juniper.net"
//
//func main() {
//	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
//	match := re.FindString(text)
//	fmt.Println(match)
//
//
//}

const text = "my email is allanyang@juniper.net " +
	"EMAIL is adv@cd.com" +
	"email is ll@qq.com " +
	" email is abd@ddd.com.cn"
const div  = "<div class=\"m-btn purple\" data-v-bff6f798=\"\">离异</div>"
const ageRe = `<div class="m-btn purple" data-v-bff6f798="">28岁</div>`
const marriageRe  = `<div class="m-btn purple" data-v-[0-9a-z]="">(离异|未婚|丧偶)</div>`
const xingzuoRe = `<div class="m-btn purple" data-v-bff6f798="">魔羯座(12.22-01.19)</div>`
const heightRe=`<div class="m-btn purple" data-v-bff6f798="">172cm</div>`
const weightRe  = `<div class="m-btn purple" data-v-bff6f798="">60kg</div>`
const hukouRe = `<div class="m-btn purple" data-v-bff6f798="">工作地:阿坝马尔康</div>`
const incomeRe  = `<div class="m-btn purple" data-v-bff6f798="">月收入:1.2-2万</div>`
const genderRe  = `<a href="http://www.zhenai.com/zhenghun/fangchenggang/nv">防城港女士征婚</a>`
const cityRe  = `<a href="http://album.zhenai.com/u/1321512066" target="_blank">一场电影而已</a>`
//func main() {
//	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
//	match := re.FindAllString(text,-1)
//	fmt.Println(match)
//
//
//}
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	//fmt.Println(match)
	for _, m := range match{
		fmt.Println(m)
	}

	re1 := regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+="">(离异|未婚|丧偶)</div>`)
	match1 := re1.FindStringSubmatch(div)
	fmt.Println(match1)

	re2 := regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(\p{Han}+)\([0-9][0-9].[0-9][0-9]-[0-9][0-9].[0-9][0-9]\)</div>`)
	match2 := re2.FindStringSubmatch(xingzuoRe)
	fmt.Println(match2)

	re3 := regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+="">([1-9][0-9]+)cm</div>`)
	match3 := re3.FindStringSubmatch(heightRe)
	fmt.Println(match3)

	re4 := regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+="">([1-9][0-9]+)kg</div>`)
	match4 := re4.FindStringSubmatch(weightRe)
	fmt.Println(match4)

	re5 := regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+="">工作地:(\p{Han}+)</div>`)
	match5 := re5.FindStringSubmatch(hukouRe)
	fmt.Println(match5)

	re6 := regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+="">月收入:([^-]*-[^-]*)</div>`)
	match6 := re6.FindStringSubmatch(incomeRe)
	fmt.Println(match6)

	re7 := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[a-z0-9]+/[a-z]+">\p{Han}+(男|女)士征婚</a>`)
	match7 := re7.FindStringSubmatch(genderRe)
	fmt.Println(match7)

	re8 := regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+="">([\d]+)岁</div>`)
	match8 := re8.FindStringSubmatch(ageRe)
	fmt.Println(match8)

	re9 := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]+>([^<]+)</a>`)
	match9 := re9.FindStringSubmatch(cityRe)
	fmt.Println(match9)
	resp, err := http.Get("http://album.zhenai.com/u/1163160974")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %v",resp)

}

