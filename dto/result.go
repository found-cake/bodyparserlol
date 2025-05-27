package dto

type PlayResult struct {
	StartTime int64 `json:"start_time"`
	Score     int   `json:"score"`
}

type GameResult struct {
	Message string `json:"msg"`
}
