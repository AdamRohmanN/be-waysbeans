package productdto

type ProductRequest struct {
	Image      string `json:"image" form:"image"`
	Name       string `json:"name" form:"name"`
	Desc       string `json:"desc" form:"desc"`
	Price      int    `json:"price" form:"price"`
	Stock      int    `json:"stock" form:"stock"`
	UserId     int    `json:"user_id"`
	CategoryId int    `json:"categories" form:"categories"`
}
