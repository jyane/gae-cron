package main

import (
	"testing"
)

var topics = []string{
	"minutely-tick",
	"hourly-tick",
	"daily-tick",
	"weekly-tick",
}

func TestStrToTopic(t *testing.T) {
	for _, v := range topics {
		_, found := strToTopic[v]
		if !found {
			t.Fatal("Not found")
		}
	}
}
