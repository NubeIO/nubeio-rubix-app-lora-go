package todos

type TODO struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	IsChecked bool `json:"is_checked"`
}
