package domain

type DefaultResponse struct {
  Status uint16 `json:"status"`
  Message string `json:"message"`
  Code string `json:"code"`
  Path string `json:"path"`
}
