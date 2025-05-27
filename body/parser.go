package body

import (
	"fmt"

	"github.com/found-cake/bodyparserlol/errors"
	"github.com/gofiber/fiber/v2"
)

// 의도한 취약점 공격 탐지 될시 CTFError를 반환 합니다.
func Parser(c *fiber.Ctx, body interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else if r == "reflect: slice index out of range" {
				err = errors.CTFError{}
			} else {
				err = fmt.Errorf("schema: panic while decoding: %v", r)
			}
		}
	}()
	err = c.BodyParser(body)
	return err
}
