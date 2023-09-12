package entities

import "time"

type GameSchedule struct {
	ScheduleID              string
	CampaignName            string
	requireReminder         bool
	schedule                string // cron schedule
	initialSessionStartTime time.Time
}
