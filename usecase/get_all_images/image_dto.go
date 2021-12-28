package get_all_images

type GetAllImagesDtoOutput struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Size         float64 `json:"size"`
	Extension    string  `json:"extension"`
	Status       string  `json:"status"`
	Url          string  `json:"url"`
	ErrorMessage string  `json:"error_message"`
}
