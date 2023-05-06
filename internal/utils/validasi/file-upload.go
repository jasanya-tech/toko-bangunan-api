package validasi

import (
	"errors"
	"mime/multipart"
	"strings"
)

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("not allowed type file. type file must be " + strings.Join(contentTypes, " "))
	} else {
		return errors.New("not found content type file")
	}
}
