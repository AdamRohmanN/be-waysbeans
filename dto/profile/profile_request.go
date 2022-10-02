package profiledto

type ProfileRequest struct {
	Image    string `json:"image" form:"image"`
	Address  string `json:"address" form:"address"`
	Postcode string `json:"postcode" form:"postcode"`
	UserId   int    `json:"user_id"`
}
