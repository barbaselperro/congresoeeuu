package base

type Senates struct {
	//	Results []struct {
	//Members []struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Api_Uri       string `json:"api_uri"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Date_of_birth string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	Party         string `json:"party"`
	Url           string `json:"url"`
	//} `json:"members"`
	//	} `json:"results"`
}

type Houses struct {
	//	Results []struct {
	//Members []struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Api_Uri       string `json:"api_uri"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Date_of_birth string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	Party         string `json:"party"`
	Url           string `json:"url"`
	//} `json:"members"`
	//	} `json:"results"`
}

type Usuario struct {
	Id   string `json:"_id"`
	Pass string `json:"password"`
	User string `json:"user"`
}
