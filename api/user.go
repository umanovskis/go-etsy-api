package api

type User struct {
	Id       int          `json:"user_id"`
	Feedback FeedbackInfo `json:"feedback_info"`
}

type FeedbackInfo struct {
	Count int
	Score int
}
