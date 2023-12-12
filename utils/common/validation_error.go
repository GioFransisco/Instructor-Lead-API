package common

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors(err error) string {
	switch err.(type) {
	case *json.UnmarshalTypeError:
		return fmt.Sprintf("%s should be a %s", err.(*json.UnmarshalTypeError).Field, err.(*json.UnmarshalTypeError).Type)
	case validator.ValidationErrors:
		castedObject, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range castedObject {
				field := err.Field()
				fieldR := []rune(field)
				fieldR[0] = unicode.ToLower(fieldR[0])

				field = string(fieldR)

				switch err.Tag() {
				case "required":
					return fmt.Sprintf("%s is required", field)
				case "email":
					return fmt.Sprintf("%s is not valid email", field)
				case "min":
					return fmt.Sprintf("%s must be more than %s character", field, err.Param())
				case "max":
					return fmt.Sprintf("%s must be %s character", field, err.Param())
				}

				break
			}
		}
	}

	return err.Error()
}

func ValidateUploadFile(file multipart.File, header *multipart.FileHeader) error {
	if file == nil {
		return fmt.Errorf("you must upload file")
	}
	
	// maxSize to upload file = 1mb
	maxSize := int64(1 << 20)
	if header.Size > maxSize {
		return fmt.Errorf("uploaded files must not exceed 1 mb")
	}
	return nil
}


