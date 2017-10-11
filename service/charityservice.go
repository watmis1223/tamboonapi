package service

type Charity struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}

func CharityList() []Charity {
	charities := []Charity{
		Charity{Id: 0, Name: "Ban Khru Noi", LogoURL: "http://rkdretailiq.com/news/img-corporate-baankrunoi.jpg"},
		Charity{Id: 1, Name: "Habitat for Humanity Thailand", LogoURL: "http://www.adamandlianne.com/uploads/2/2/1/6/2216267/3231127.gif"},
		Charity{Id: 2, Name: "Paper Ranger", LogoURL: "https://myfreezer.files.wordpress.com/2007/06/paperranger.jpg"},
		Charity{Id: 3, Name: "Makhampom", LogoURL: "http://www.makhampom.net/makhampom/ppcms/uploads/UserFiles/Image/Thai/T14Publice/2554/January/Newyear/logoweb.jpg"},
	}
    return charities
}