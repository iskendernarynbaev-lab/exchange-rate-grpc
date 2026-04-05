package model

import "time"

type Rate struct {
	Ask       float64
	Bid       float64
	CreatedAt time.Time
}
