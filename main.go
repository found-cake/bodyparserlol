package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/found-cake/bodyparserlol/body"
	"github.com/found-cake/bodyparserlol/config"
	"github.com/found-cake/bodyparserlol/dto"
	"github.com/found-cake/bodyparserlol/errors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/random", randomDice)
	app.Post("/now", getNow)
	app.Post("/result", checkClear)

	log.Fatal(app.Listen(":3000"))
}

func getNow(c *fiber.Ctx) error {
	now := time.Now().Unix()
	return c.JSON(map[string]int64{
		"time": now,
	})
}

func checkClear(c *fiber.Ctx) error {
	result := dto.PlayResult{}
	if err := body.Parser(c, &result); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if time.Now().Unix()-result.StartTime > int64(config.MaxPlayTime) {
		return c.JSON(dto.GameResult{
			Message: "플레이 타임 초과 되었습니다.",
		})
	}
	if result.Score != config.RequiredScore {
		return c.JSON(dto.GameResult{
			Message: fmt.Sprintf("목표 점수랑 %d 점 차이가 있습니다.", result.Score-config.RequiredScore),
		})
	}
	return c.JSON(dto.GameResult{
		Message: fmt.Sprintf("Flag is %s  ㅋㅋ", config.FAKE_FLAG),
	})
}

func randomDice(c *fiber.Ctx) error {
	info := dto.Selector{}
	if err := body.Parser(c, &info); err != nil {
		if (err == errors.CTFError{}) {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		log.Println(err)
		return c.SendStatus(http.StatusBadRequest)
	}
	total := 0
	randCount := 0
	for _, v := range info.Dice {
		if v == nil {
			continue
		}
		for i := 0; i < v.Count; i++ {
			total += 1 + rand.Intn(v.DiceType)
			randCount++
		}
	}
	if randCount > config.MaxRandCount {
		return c.JSON(dto.RandResult{
			Message: fmt.Sprintf("주사위 최대 겟수는 %d개 입니다.", config.MaxRandCount),
			Total:   0,
		})
	}
	return c.JSON(dto.RandResult{
		Message: fmt.Sprintf("%d점 휙득", total),
		Total:   total,
	})
}
