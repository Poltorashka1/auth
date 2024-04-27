package response

type Response interface {
	StatusCode() int
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func Error(err error, status int) Response {
	return ErrorResponse{
		Error:  err.Error(),
		Status: status,
	}
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

type SuccessResponse struct {
	Status int         `json:"status"`
	Type   string      `json:"type,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func (s SuccessResponse) StatusCode() int {
	return s.Status
}

func Success(data interface{}, typ ...string) Response {
	if data == nil {
		return SuccessResponse{
			Status: 200,
		}
	}

	res := SuccessResponse{
		Status: 200,
		Result: data,
	}
	if typ != nil {
		res.Type = typ[0]
	}
	return res
}
