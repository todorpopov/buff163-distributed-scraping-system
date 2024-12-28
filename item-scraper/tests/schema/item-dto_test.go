package schema_test

import (
	"strings"
	"testing"

	"github.com/todorpopov/bdss-item-scraper/src/schema"
)

func TestSerialize(t *testing.T) {
	stickers := []schema.StickerDTO{
		schema.StickerDTO{
			Name: "name",
			Slot: 1,
			Wear: 1,
		},
		schema.StickerDTO{
			Name: "name2",
			Slot: 2,
			Wear: 2,
		},
	}

	dto := schema.ItemDTO{
		UserId:               "test",
		Price:                "100",
		ImgSrc:               "URL",
		AssetId:              "id",
		HasTradeableCooldown: true,
		Id:                   "id",
		PaintIndex:           1,
		PaintSeed:            1,
		Stickers:             stickers,
		OfferUrl:             "url",
	}

	expected := "{\"UserId\":\"test\",\"Price\":\"100\",\"ImgSrc\":\"URL\",\"AssetId\":\"id\",\"HasTradeableCooldown\":true,\"Id\":\"id\",\"PaintIndex\":1,\"PaintSeed\":1,\"Stickers\":[{\"Name\":\"name\",\"Slot\":1,\"Wear\":1},{\"Name\":\"name2\",\"Slot\":2,\"Wear\":2}],\"OfferUrl\":\"url\"}"
	actual := dto.Serialize()

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Actual result: %s", actual)
	}
}
