package constant

// HttpMethod HTTP 请求方式
type HttpMethod string

const (
	Get     HttpMethod = "GET"
	Head    HttpMethod = "HEAD"
	Post    HttpMethod = "POST"
	Put     HttpMethod = "PUT"
	Patch   HttpMethod = "PATCH" // RFC 5789
	Delete  HttpMethod = "DELETE"
	Connect HttpMethod = "CONNECT"
	Options HttpMethod = "OPTIONS"
	Trace   HttpMethod = "TRACE"
)

func (h HttpMethod) Value() string {
	return string(h)
}
