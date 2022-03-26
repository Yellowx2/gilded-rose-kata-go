package main

import (
	"testing"
)

func TestLowerSellInDate(t *testing.T) {
	tests := []struct {
		name string
		item *Item
	}{
		// TODO: Add test cases.
		{"happy path", &Item{"foo", 0, 0}},
		{"another happy path", &Item{"bar", 10, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalItem := *tt.item
			UpdateQuality([]*Item{tt.item})
			// Only lowered by one
			if originalItem.sellIn-tt.item.sellIn != 1 {
				t.Errorf("Item %s: want %v, got %v", tt.item.name, originalItem.sellIn-1, tt.item.sellIn)
			}
		})
	}
}
