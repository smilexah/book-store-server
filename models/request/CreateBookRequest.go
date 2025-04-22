package request

type CreateBookRequest struct {
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
