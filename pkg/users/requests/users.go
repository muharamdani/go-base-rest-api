package requests

type GetUser struct {
	Limit int `form:"limit,default=10"`
	Page  int `form:"page,default=1"`
}
