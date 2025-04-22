package request

type UpdateBookRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
