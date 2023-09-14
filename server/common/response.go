package common

type HttpResponse struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type GRPCResponse struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
