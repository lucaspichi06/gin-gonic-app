package domain

type User struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	UserName string  `json:"username"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
	Phone   string  `json:"phone"`
	WebSite string  `json:"website"`
	Company Company `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}
