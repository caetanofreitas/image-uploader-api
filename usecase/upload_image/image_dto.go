package upload_image

type UploadImageDtoInput struct {
	ID        string
	Name      string
	Size      float64
	Extension string
	Image     []byte
}

type UploadImageDtoOutput struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
