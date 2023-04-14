package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"toko-bangunan/helpers/exception"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFileImg(c *fiber.Ctx) error {
	file, _ := c.FormFile("image")
	var format string
	if file != nil {
		if err := checkContentType(file, "image/png", "image/jpeg", "image/jpg"); err != nil {
			panic(exception.BadRequestError{Message: err.Error()})
		}
		format = fmt.Sprintf("%s%s%s", time.Time.Format(time.Now(), "06010205"), time.Nanosecond.String(), filepath.Ext(file.Filename))
	}

	if format != "nil" {
		c.Locals("filename", format)
	} else {
		c.Locals("filename", nil)
	}

	return c.Next()
}

func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
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
