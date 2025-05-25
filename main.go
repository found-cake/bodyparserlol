package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CTFError

const FLAG = "SF{...}"

type Selector struct {
	Dice []*struct {
		DiceType int `form:"type"`
		Count    int `form:"count"`
	} `form:"dice_info"`
}

func main() {
	app := fiber.New()

	app.Post("/random", randomDice)

	log.Fatal(app.Listen(":3000"))
}

func parser(c *fiber.Ctx, body interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else if r == "reflect: slice index out of range" {
				//TODO
				err = fmt.Errorf(FLAG)
			} else {
				err = fmt.Errorf("schema: panic while decoding: %v", r)
			}
		}
	}()

	err = c.BodyParser(body)

	return err
}

func randomDice(c *fiber.Ctx) error {
	info := Selector{}
	if err := parser(c, &info); err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusBadRequest)
	}
	total := 0
	for _, v := range info.Dice {
		if v == nil {
			continue
		}
		for i := 0; i < v.Count; i++ {
			total += 1 + rand.Intn(v.DiceType)
		}
	}
	return c.JSON(map[string]int{
		"total": total,
	})
}
