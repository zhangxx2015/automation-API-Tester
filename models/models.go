package models

type RequestInfo struct {
	Url     string                 `json:"url"`
	Querys  map[string]interface{} `json:"querys"`
	Headers map[string]interface{} `json:"headers"`
	Method  string                 `json:"method"`
	Body    map[string]interface{} `json:"body"`
	Assert  map[string]interface{} `json:"assert"`
}
