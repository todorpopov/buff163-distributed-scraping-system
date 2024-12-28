package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetItemsApi(itemCode string) string {
	return fmt.Sprintf("https://buff.163.com/api/market/goods/sell_order?game=csgo&goods_id=%s&page_num=1", itemCode)
}

func GetItemOfferUrl(userId string, itemName string) string {
	escapedItemName := strings.ReplaceAll(itemName, " ", "%20")
	return fmt.Sprintf("https://buff.163.com/shop/%s#tab=selling&game=csgo&page_num=1&search=%s", userId, escapedItemName)
}

func ParseItemCodesFile() []string {
	data, err := os.ReadFile("./item-codes.txt")
	if err != nil {
		fmt.Printf("Error occured: %s", err)
	}

	return strings.Split(string(data), "|")
}
