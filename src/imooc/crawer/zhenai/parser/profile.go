package parser

import (
	"imooc/crawer/engine"
	"imooc/crawer/model"
	"regexp"
	"strconv"
)

var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[a-z0-9]+/[a-z]+">\p{Han}+(男|女)士征婚</a>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+>([\d]+)岁</div>`)

var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+>(离异|未婚|丧偶)</div>`)
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+>(\p{Han}+)\([0-9][0-9].[0-9][0-9]-[0-9][0-9].[0-9][0-9]\)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+>([1-9][0-9]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+>([1-9][0-9]+)kg</div>`)
var hukouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+>工作地:(\p{Han}+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-[0-9a-z]+>月收入:([^-]*-[^-]*)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	//fmt.Printf("age: %v",age)
	if err != nil {
		if err != nil {
			//fmt.Errorf("Did not get age value:%s", err)
			profile.Age = 0
		}
	}
	profile.Age = age
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		//fmt.Errorf("Did not get height value:%s", err)
		profile.Height = 0
	}
	profile.Height = height
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		//fmt.Errorf("Did not get weight value:%s", err)
		profile.Weight = 0
	}
	profile.Weight = weight
	profile.Gender = extractString(contents, genderRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(context []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(context)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
