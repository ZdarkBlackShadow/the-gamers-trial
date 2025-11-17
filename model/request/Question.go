package request

type Question struct {
	QuestionID     int `json:"question_id"`
	UserAnswerID   int `json:"user_answer_id"`
	ResponseTimeMs int `json:"response_time_ms"`
}
