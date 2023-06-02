package mybots

//RequestData data struct for json
type RequestData struct {
	TerminalId string `json:"terminal_id"`
	Query      string `json:"query"`
}

// ConversationRequest data struct for json
type ConversationRequest struct {
	Version   string      `json:"version"`
	ServiceId string      `json:"service_id"`
	LogId     string      `json:"log_id"`
	SessionId string      `json:"session_id"`
	Request   RequestData `json:"request"`
}

func newConversationRequest(request string, sessionId string, terminalId string) ConversationRequest {
	requestData := RequestData{TerminalId: terminalId, Query: request}
	return ConversationRequest{Version: "3.0", ServiceId: "S90926", LogId: terminalId, SessionId: sessionId, Request: requestData}
}
