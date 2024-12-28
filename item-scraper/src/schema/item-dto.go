package schema

import (
	"encoding/json"
)

type StickerDTO struct {
	Name string
	Slot float64
	Wear float64
}

type ItemDTO struct {
	UserId               string
	Price                string
	ImgSrc               string
	AssetId              string
	HasTradeableCooldown bool
	Id                   string
	PaintIndex           float64
	PaintSeed            float64
	Stickers             []StickerDTO
	OfferUrl             string
}

func (dto *ItemDTO) Serialize() string {
	json, err := json.Marshal(dto)
	if err != nil {
		return ""
	}

	return string(json)
}
