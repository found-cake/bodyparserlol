package dto

type Selector struct {
	Dice []*struct {
		DiceType int `form:"type"`
		Count    int `form:"count"`
	} `form:"dice_info"`
}
