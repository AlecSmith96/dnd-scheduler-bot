package entities

import "time"

type GameSchedule struct {
	CampaignName            string
	requireReminder         bool
	schedule                string // cron schedule
	initialSessionStartTime time.Time
}
