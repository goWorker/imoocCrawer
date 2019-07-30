package parser

import (
	"imooc/crawer/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contests,err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contests,"阿坝吃虫草的牦牛")
	if len(result.Items) != 1{
		t.Errorf("Items should contain 1 element; but was %v",result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:			"阿坝吃虫草的牦牛",
		Gender: 		"男",
		Marriage:	 	"离异",
		Age:	 		34,
		Xingzuo:		"魔羯座",
		Height:			172,
		Weight: 		60,
		Hukou:			"阿坝马尔康",
		Income:			"1.2-2万",
	}

	if profile != expected {
		t.Errorf("Expected %v, but was %v",expected,profile)
	}
}
