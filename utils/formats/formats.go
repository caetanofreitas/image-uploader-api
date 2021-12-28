package formats

import "errors"

const (
	JPG  = "jpg"
	JPEG = "jpeg"
	JFIF = "jfif"
)

func ValidateFormat(format string) error {
	if format != JPG &&
		format != JPEG &&
		format != JFIF {
		return errors.New("invalid")
	}
	return nil
}
