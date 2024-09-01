package domain


type OKResponse struct {
  Status uint16 `json:"status"`
  Message string `json:"message"`
  Code string `json:"code"`
  Data any `json:"data"`
}