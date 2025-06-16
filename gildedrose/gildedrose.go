package gildedrose

import (
	"strings"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}
		items[i].SellIn = items[i].SellIn - 1

		if items[i].Name == "Aged Brie" {
			items[i].Quality = items[i].Quality + 1
			if items[i].SellIn < 0 {
				items[i].Quality = items[i].Quality + 1
			}
		} else if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].SellIn < 0 {
				items[i].Quality = 0
			} else if items[i].SellIn < 5 {
				items[i].Quality = items[i].Quality + 3
			} else if items[i].SellIn < 10 {
				items[i].Quality = items[i].Quality + 2
			} else {
				items[i].Quality = items[i].Quality + 1
			}
		} else if strings.HasPrefix(items[i].Name, "Conjured") {
			items[i].Quality = items[i].Quality - 2
			if items[i].SellIn < 0 {
				items[i].Quality = items[i].Quality - 2
			}
		} else {
			items[i].Quality = items[i].Quality - 1
			if items[i].SellIn < 0 {
				items[i].Quality = items[i].Quality - 1
			}
		}

		if items[i].Quality < 0 {
			items[i].Quality = 0
		}
		if items[i].Quality > 50 {
			items[i].Quality = 50
		}
	}
}
