package domain

type NotFoundResponse struct {
  Status uint16 `json:"status"`
  Message string `json:"message"`
  Code string `json:"code"`
}
