package models

type Category struct {
	Id 		int
	Pid 	int
	Title 	string `orm:"size(30);unique"`
	Intro 	string
	Icon 	string
	Cnt 	int   //统计分类图书
	Sort 	int
	Status 	bool
}

func (m *Category)TableName() string{
	return TNCategory()
}