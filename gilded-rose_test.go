package main

import (
	"reflect"
	"testing"
)

func TestLowerSellInDate(t *testing.T) {
	tests := []struct {
		name string
		item *Item
		want *Item
	}{
		// TODO: Add test cases.
		{"happy path", &Item{"bar", 10, 6}, &Item{"bar", 9, 5}},
		{"happy path with 0 quality", &Item{"foo", 0, 0}, &Item{"foo", -1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reflect.DeepEqual(*tt.item, *tt.want) {
				t.Errorf("Got -> %v\n Want -> %v", *tt.item, tt.want)
			}
		})
	}
}
