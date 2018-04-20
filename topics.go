package main

type Topic string

const (
	MinutelyTick Topic = "minutely-tick"
	HourlyTick   Topic = "hourly-tick"
	DailyTick    Topic = "daily-tick"
	WeeklyTick   Topic = "weekly-tick"
)

var strToTopic = map[string]Topic{
	"minutely-tick": MinutelyTick,
	"hourly-tick":   HourlyTick,
	"daily-tick":    DailyTick,
	"weekly-tick":   WeeklyTick,
}
