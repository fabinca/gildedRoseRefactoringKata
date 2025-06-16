package gildedrose_test

import (
	"testing"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_NormalItem(t *testing.T) {
	var items = []*gildedrose.Item{
		{"rose", 1, 4},
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 0 || items[0].Quality != 3 {
		t.Errorf("Name: Expected {'rose', 0, 3} but got {'rose', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != -1 || items[0].Quality != 1 {
		t.Errorf("Name: Expected {'rose', -1, 1} but got {'rose', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != -2 || items[0].Quality != 0 {
		t.Errorf("Name: Expected {'rose', -2, 0} but got {'rose', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}
}

func Test_AgedBrie(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 1, 46},
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 0 || items[0].Quality != 47 {
		t.Errorf("Name: Expected {'Aged Brie', 0, 47} but got {'Aged Brie', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != -1 || items[0].Quality != 49 {
		t.Errorf("Name: Expected {'Aged Brie', -1, 49} but got {'Aged Brie', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != -2 || items[0].Quality != 50 {
		t.Errorf("Name: Expected {'Aged Brie', -2, 50} but got {'Aged Brie', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}
}

func Test_Sulfuras(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 1, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 25},
		{"Sulfuras, Hand of Ragnaros", 5, 3},
		{"Sulfuras, Hand of Ragnaros", 4, 55},
	}

	gildedrose.UpdateQuality(items)


	if items[0].SellIn != 1 || items[0].Quality != 80 {
		t.Errorf("Name: Expected {'Sulfuras', 1, 80} but got {'Sulfuras', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}
	if items[1].SellIn != -1 || items[1].Quality != 25 {
		t.Errorf("Name: Expected {'Sulfuras', -1, 25} but got {'Sulfuras', %d, %d} ", items[1].SellIn, items[1].Quality, )
	}
	if items[2].SellIn != 5 || items[2].Quality != 3 {
		t.Errorf("Name: Expected {'Sulfuras', 5, 3} but got {'Sulfuras', %d, %d} ", items[2].SellIn, items[2].Quality, )
	}
	if items[3].SellIn != 4 || items[3].Quality != 55 {
		t.Errorf("Name: Expected {'Sulfuras', 4, 55} but got {'Sulfuras', %d, %d} ", items[3].SellIn, items[3].Quality, )
	}
}

func Test_BackstagePasses(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 11, 3},
		{"Backstage passes to a TAFKAL80ETC concert", 6, 3},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 3},
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 14 || items[0].Quality != 50 {
		t.Errorf("Name: Expected {'BackstagePass', 14, 50} but got {'BackstagePass', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}
	if items[1].SellIn != 10 || items[1].Quality != 4 {
		t.Errorf("Name: Expected {'BackstagePass', 10, 4} but got {'BackstagePass', %d, %d} ", items[1].SellIn, items[1].Quality, )
	}
	if items[2].SellIn != 5 || items[2].Quality != 5 {
		t.Errorf("Name: Expected {'BackstagePass', 5, 5} but got {'BackstagePass', %d, %d} ", items[2].SellIn, items[2].Quality, )
	}
	if items[3].SellIn != 0 || items[3].Quality != 6 {
		t.Errorf("Name: Expected {'BackstagePass', 0, 6} but got {'BackstagePass', %d, %d} ", items[3].SellIn, items[3].Quality, )
	}

	gildedrose.UpdateQuality(items)

		if items[0].SellIn != 13 || items[0].Quality != 50 {
		t.Errorf("Name: Expected {'BackstagePass', 13, 50} but got {'BackstagePass', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}
	if items[1].SellIn != 9 || items[1].Quality != 6 {
		t.Errorf("Name: Expected {'BackstagePass', 9, 6} but got {'BackstagePass', %d, %d} ", items[1].SellIn, items[1].Quality, )
	}
	if items[2].SellIn != 4 || items[2].Quality != 8 {
		t.Errorf("Name: Expected {'BackstagePass', 4, 8} but got {'BackstagePass', %d, %d} ", items[2].SellIn, items[2].Quality, )
	}
	if items[3].SellIn != -1 || items[3].Quality != 0 {
		t.Errorf("Name: Expected {'BackstagePass', -1, 0} but got {'BackstagePass', %d, %d} ", items[3].SellIn, items[3].Quality, )
	}
}

func Test_ConjuredItem(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 1, 7},
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 0 || items[0].Quality != 5 {
		t.Errorf("Name: Expected {'Conjured', 0, 5} but got {'Conjured', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != -1 || items[0].Quality != 1 {
		t.Errorf("Name: Expected {'Conjured', -1, 1} but got {'Conjured', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != -2 || items[0].Quality != 0 {
		t.Errorf("Name: Expected {'Conjured', -2, 0} but got {'Conjured', %d, %d} ", items[0].SellIn, items[0].Quality, )
	}
}
