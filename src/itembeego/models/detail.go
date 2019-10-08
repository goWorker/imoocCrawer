package models

import "fmt"

type photo struct {
	Id        string
	ImagePath string
}

func GetPhotos() []photo {
	buffer := []photo{
		photo{
			Id:        "pic1",
			ImagePath: "img01.png",
		},
		photo{
			Id:        "pic2",
			ImagePath: "img02.png",
		},
		photo{
			Id:        "pic3",
			ImagePath: "img03.png",
		},
		photo{
			Id:        "pic4",
			ImagePath: "img04.png",
		},
		photo{
			Id:        "pic5",
			ImagePath: "img05.png",
		},
		photo{
			Id:        "pic6",
			ImagePath: "img06.png",
		},
	}
	return buffer
}

func GetTitle() string {
	return "瞎编的价格"
}

func GetPriceString() string {
	price := 139900
	intPrice := price / 100
	decPrice := price % 100

	return fmt.Sprintf("%d.%02d", intPrice, decPrice)
}
