package helper

type JSONResponse struct {
	Status  bool     `json:"status"`
	Massage string   `json:"message"`
	Errors  []string `json:"errors"`
	Data    any      `json:"data"`
}
