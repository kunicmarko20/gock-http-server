package mock

type Response struct {
	StatusCode int
	Headers    map[string]string
	Content    string
}
