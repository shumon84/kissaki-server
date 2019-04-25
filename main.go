package main

import (
	"time"
)

type StageMap struct {
	Data [16]uint16 `json:"data"`
}

type StageInfo struct {
	Name          string    `json:"name"`
	StageMap      StageMap  `json:"stage_map"`
	StartPosition uint8     `json:"start_position"`
	GoalPosition  uint8     `json:"goal_position"`
	CreatedAt     time.Time `json:"created_at"`
}
