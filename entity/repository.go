package entity

type ImageRepository interface {
	Insert(id string, name string, size float64, extension string, status string, error_message string) error
	Update(id string, name string, size float64, extension string, status string, error_message string) error
	Get() ([]Image, error)
	GetDetail(id string) (Image, error)
}

type ImageUploader interface {
	UploadImage(image []byte, id string, extension string) (int64, error)
}
