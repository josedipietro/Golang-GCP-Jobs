package structs

import (
	"time"
)

type Job struct {
	ID        string `json:"id"`
	Desc      string `json:"desc"`
	JobType   int64  `json:"jobType"`
	Priority  int64  `json:"priority"`
	ExecType  int64  `json:"execType"`
	Payload   Payload  `json:"payload"`
	Args      interface{}  `json:"args"`
	CreatedAt time.Time `json:"createdAt"`
}
