package request

type CreateBookReqBody struct {
	Title       string `json:"title" binding:"required,min=3,max=25"`
	AuthorName  string `json:"author_name" binding:"required,min=3,max=25"`
	Publication string `json:"publication" binding:"required,min=3,max=25"`
}

type UpdateBookReqBody struct {
	Title       string `json:"title" binding:"min=3,max=25"`
	AuthorName  string `json:"author_name" binding:"min=3,max=25"`
	Publication string `json:"publication" binding:"min=3,max=25"`
}
