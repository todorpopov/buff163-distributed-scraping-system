package schema

type Sticker struct {
	Name string `json:"name"`
	Slot int    `json:"slot"`
	Wear int    `json:"wear"`
}

type Info struct {
	PaintIndex int       `json:"paintindex"`
	PaintSeed  int       `json:"paintseed"`
	Stickers   []Sticker `json:"stickers"`
}

type AssetInfo struct {
	HasTradeableCooldown bool   `json:"has_tradable_cooldown"`
	Id                   string `json:"id"`
	Info                 Info   `json:"info"`
}

type Item struct {
	AssetInfo AssetInfo `json:"asset_info"`
	ImgSrc    string    `json:"img_src"`
	Price     string    `json:"price"`
	UserId    string    `json:"user_id"`
}

type Data struct {
	Items []Item `json:"items"`
}

type ResponseJson struct {
	Data Data `json:"data"`
}

func (r ResponseJson) ToItemDtos() *[]ItemDTO {
	var itemDtos []ItemDTO

	for i := 0; i < len(r.Data.Items); i++ {
		var item ItemDTO

		item.UserId = r.Data.Items[i].UserId
		item.Price = r.Data.Items[i].Price
		item.ImgSrc = r.Data.Items[i].ImgSrc
		item.Id = r.Data.Items[i].AssetInfo.Id
		item.HasTradeableCooldown = r.Data.Items[i].AssetInfo.HasTradeableCooldown
		item.PaintIndex = r.Data.Items[i].AssetInfo.Info.PaintIndex
		item.PaintSeed = r.Data.Items[i].AssetInfo.Info.PaintSeed

		var stickers []StickerDTO
		for j := 0; j < len(r.Data.Items[i].AssetInfo.Info.Stickers); j++ {
			var sticker StickerDTO

			sticker.Name = r.Data.Items[i].AssetInfo.Info.Stickers[j].Name
			sticker.Slot = r.Data.Items[i].AssetInfo.Info.Stickers[j].Slot
			sticker.Wear = r.Data.Items[i].AssetInfo.Info.Stickers[j].Wear

			stickers = append(stickers, sticker)
		}

		item.Stickers = stickers

		itemDtos = append(itemDtos, item)
	}

	return &itemDtos
}
