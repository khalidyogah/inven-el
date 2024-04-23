package structs

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       string `json:"price"`
	TotalPage   int    `json:"total_page"`
	Thickness   string `json:"thickness"`
	CreatedAt   string
	UpdatedAt   string
	CategoryId  int `json:"category_id"`
}

func (s *Book) DecideThickness() {
	if s.TotalPage <= 100 {
		s.Thickness = "tipis"
	} else if s.TotalPage > 100 && s.TotalPage < 200 {
		s.Thickness = "sedang"
	} else {
		s.Thickness = "tebal"
	}
}
