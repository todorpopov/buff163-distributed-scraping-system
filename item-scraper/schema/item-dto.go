package schema

type StickerDTO struct {
	Name string
	Slot int
	Wear int
}

type ItemDTO struct {
	UserId               string
	Price                string
	ImgSrc               string
	AssetId              string
	HasTradeableCooldown bool
	Id                   string
	PaintIndex           int
	PaintSeed            int
	Stickers             []StickerDTO
	OfferUrl             string
}
