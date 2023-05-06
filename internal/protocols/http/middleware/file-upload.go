package middleware

import (
	"fmt"
	"path/filepath"
	"time"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/internal/utils/validasi"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFileImg(c *fiber.Ctx) error {
	file, _ := c.FormFile("image")
	var format string
	if file != nil {
		if err := validasi.CheckContentType(file, "image/png", "image/jpeg", "image/jpg"); err != nil {
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
