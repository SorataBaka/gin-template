package domain



type ErrorResponse struct {
  Status uint16 `json:"status"`
  Message string `json:"message"`
  Code string `json:"code"`
}