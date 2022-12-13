package response

type Result struct {
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (r *Result) Success() (Status int, Data any) {
	return r.Status, r.Data
}

func (r *Result) Error() (Status int, Code string, Message string) {
	return r.Status, r.Code, r.Message
}
