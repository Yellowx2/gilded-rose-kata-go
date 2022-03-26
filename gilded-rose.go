package main

type Item struct {
	name            string
	sellIn, quality int
}

const (
	AgedBrie        = "Aged Brie"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Conjured        = "Conjured"
)

// func UpdateQuality(items []*Item) {
// 	for i := 0; i < len(items); i++ {

// 		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
// 			if items[i].quality > 0 {
// 				if items[i].name != "Sulfuras, Hand of Ragnaros" {
// 					items[i].quality = items[i].quality - 1
// 				}
// 				if items[i].name == "Conjured" && items[i].quality > 0 {
// 					items[i].quality = items[i].quality - 1
// 				}
// 			}
// 		} else {
// 			if items[i].quality < 50 {
// 				items[i].quality = items[i].quality + 1
// 				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
// 					if items[i].sellIn < 11 {
// 						if items[i].quality < 50 {
// 							items[i].quality = items[i].quality + 1
// 						}
// 					}
// 					if items[i].sellIn < 6 {
// 						if items[i].quality < 50 {
// 							items[i].quality = items[i].quality + 1
// 						}
// 					}
// 				}
// 			}
// 		}

// 		if items[i].name != "Sulfuras, Hand of Ragnaros" {
// 			items[i].sellIn = items[i].sellIn - 1
// 		}

// 		if items[i].sellIn < 0 {
// 			if items[i].name != "Aged Brie" {
// 				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
// 					if items[i].quality > 0 {
// 						if items[i].name != "Sulfuras, Hand of Ragnaros" {
// 							items[i].quality = items[i].quality - 1
// 						}
// 						if items[i].name == "Conjured" && items[i].quality > 0 {
// 							items[i].quality = items[i].quality - 1
// 						}
// 					}
// 				} else {
// 					items[i].quality = items[i].quality - items[i].quality
// 				}
// 			} else {
// 				if items[i].quality < 50 {
// 					items[i].quality = items[i].quality + 1
// 				}
// 			}
// 		}
// 	}

// }

func UpdateQuality(items []*Item) {

	for _, item := range items {

		switch item.name {

		case AgedBrie:
			decreaseSellIn(item)
			increaseQuality(item)
			if item.sellIn < 0 {
				increaseQuality(item)
			}

		case BackstagePasses:
			decreaseSellIn(item)
			if item.sellIn >= 0 {
				increaseQuality(item)
				if item.sellIn <= 10 {
					increaseQuality(item)
					if item.sellIn <= 5 {
						increaseQuality(item)
					}
				}
			} else {
				if item.quality != 0 {
					item.quality = 0
				}
			}

		case Sulfuras:
			continue

		case Conjured:
			decreaseSellIn(item)
			decreaseQuality(item)
			decreaseQuality(item)

		default:
			decreaseSellIn(item)
			decreaseQuality(item)
		}
	}
}

func decreaseQuality(item *Item) {
	if item.quality > 0 {
		item.quality--
		if item.sellIn < 0 && item.quality > 0 {
			item.quality--
		}
	}
}

func increaseQuality(item *Item) {
	if item.quality < 50 {
		item.quality++
	}
}

func decreaseSellIn(item *Item) {
	item.sellIn--
}
