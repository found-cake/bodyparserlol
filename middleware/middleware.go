package middleware

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

var re = regexp.MustCompile(`\[\-+\d`)

func Filter(c *fiber.Ctx) error {
	b := c.Body()
	b = re.ReplaceAll(b, []byte("["))
	c.Request().SetBody(b)
	c.Request().Header.SetContentLength(len(b))
	return c.Next()
}
