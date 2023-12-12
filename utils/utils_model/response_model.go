package utilsmodel

type Response struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type ResponseError struct {
	Response Response `json:"response"`
}

type ResponseSuccess struct {
	Response Response `json:"response"`
	Data     any      `json:"data"`
}
