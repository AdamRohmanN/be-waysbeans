package authdto

type LoginResponse struct {
	Status   string `json:"status"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type CheckAuthResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
