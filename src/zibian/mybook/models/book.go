package models


type Book struct{
	BookId 		int `orm:"pk;auto" json:"book_id"`
	BookName 	string `orm:"size(500)" json:"book_name"`
	Identify 	string `orm:"size(100);unique" json:"identify"`
	OrderIndex 	int `orm:"default(0)" json:"order_index"`
	Description string `orm:"size(1000)" json:"description"`
}

func (m *Book)TableName()string{
	return TNBook()
}