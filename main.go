package main

import (
	"encoding/json"
	"log"
	"net/http"
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

var (
	MaxNumOfStages = 30
	DB             = make([]StageInfo, 0)
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	stageInfo := new(StageInfo)
	if err := json.NewDecoder(r.Body).Decode(stageInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	stageInfo.CreatedAt = time.Now()
	if len(DB) >= MaxNumOfStages {
		DB = DB[1:MaxNumOfStages]
	}
	DB = append(DB, *stageInfo)
}

func GetHandler(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(DB)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetHandler(w, r)
		case http.MethodPost:
			PostHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	})))
}
