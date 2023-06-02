package mybots

//ConversionAction represents
type ConversionAction struct {
	Confidence float64     `json:"confidence"`
	Say        string      `json:"say"`
	Type       string      `json:"type"`
	ActionId   string      `json:"action_id"`
	Img        interface{} `json:"img"`
}

//ConversationsResponse represents
type ConversationsResponse struct {
	Status      int                `json:"status"`
	Msg         string             `json:"msg"`
	Origin      string             `json:"origin"`
	Schema      interface{}        `json:"schema"`
	Actions     []ConversionAction `json:"actions"`
	RawQuery    string             `json:"raw_query"`
	SlotHistory interface{}        `json:"slot_history"`
}

//ConversationResult represents
type ConversationResult struct {
	Version   string                  `json:"version"`
	Context   interface{}             `json:"context"`
	TimeStamp string                  `json:"timestamp"`
	ServiceId string                  `json:"service_id"`
	SessionId string                  `json:"session_id"`
	LogId     string                  `json:"log_id"`
	RefId     string                  `json:"ref_id"`
	Responses []ConversationsResponse `json:"responses"`
}

//ConversationResponse represents
type ConversationResponse struct {
	ErrorCode int                `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
	Result    ConversationResult `json:"result"`
}
