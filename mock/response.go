package mock

type Response struct {
	statusCode int
	headers    map[string]string
	body       string
}

func NewResponse(statusCode int, headers map[string]string, body string) *Response {
	return &Response{statusCode, headers, body}
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Headers() map[string]string {
	return r.headers
}

func (r *Response) Body() string {
	return r.body
}
