package main

import (
	"reflect"
	"testing"
)

func TestUpdateQuality(t *testing.T) {
	tests := []struct {
		name string
		item *Item
		want *Item
	}{
		{"system lowers both values", &Item{"bar", 10, 6}, &Item{"bar", 9, 5}},
		{"when sell in date below 0, quality degrades by 2", &Item{"foo", 0, 10}, &Item{"foo", -1, 8}},
		{"quality can't be below 0", &Item{"foo", 0, 0}, &Item{"foo", -1, 0}},
		{"Aged Brie increases quality by 1", &Item{"Aged Brie", 25, 9}, &Item{"Aged Brie", 24, 10}},
		{"Aged Brie increases quality by 2 after sell in date below 0", &Item{"Aged Brie", 0, 9}, &Item{"Aged Brie", -1, 11}},
		{"quality of an item is never more than 50", &Item{"Aged Brie", 0, 50}, &Item{"Aged Brie", -1, 50}},
		{"quality of an item is never more than 50", &Item{"Aged Brie", 0, 49}, &Item{"Aged Brie", -1, 50}},
		{"'Sulfuras' never has to be sold or decreases in Quality", &Item{"Sulfuras, Hand of Ragnaros", 0, 80}, &Item{"Sulfuras, Hand of Ragnaros", 0, 80}},
		{"'Backstage passes' quality increases by 2 when there are 10 days or less", &Item{"Backstage passes to a TAFKAL80ETC concert", 10, 5}, &Item{"Backstage passes to a TAFKAL80ETC concert", 9, 7}},
		{"'Backstage passes' quality increases by by 3 when there are 5 days or less", &Item{"Backstage passes to a TAFKAL80ETC concert", 5, 6}, &Item{"Backstage passes to a TAFKAL80ETC concert", 4, 9}},
		{"'Backstage passes' quality drops to 0 after the concert", &Item{"Backstage passes to a TAFKAL80ETC concert", 0, 15}, &Item{"Backstage passes to a TAFKAL80ETC concert", -1, 0}},
		{"'Conjured' items degrade in quality twice as fast as normal items", &Item{"Conjured", 15, 10}, &Item{"Conjured", 14, 8}},
		{"'Conjured' items degrade in quality twice as fast after sell in date below 0", &Item{"Conjured", 0, 11}, &Item{"Conjured", -1, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality([]*Item{tt.item})
			if !reflect.DeepEqual(tt.item, tt.want) {
				t.Errorf("Want -> %v\nGot -> %v", *tt.want, *tt.item)
			}
		})
	}
}
